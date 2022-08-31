package scalingo

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/Scalingo/go-scalingo/v4/http/httpmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var isDeprecatedCases = map[string]struct {
	date       time.Time
	deprecated bool
}{
	"test isDeprecated is false when deprecation date is null|nil|zero": {time.Time{}, false},
	"test isDeprecated is true when deprecation date is today's date":   {time.Now(), true},
	"test isDeprecated is true when deprecation date is in the past":    {time.Now().AddDate(0, 0, -1), true},
	"test isDeprecated is false when deprecation date is in the future": {time.Now().AddDate(0, 0, 1), false},
}

var stacksListCases = map[string]struct {
	json          string
	expectedStack Stack
}{
	"test StacksList with stacks with deprecated_at being null": {
		json:          `{"stacks": [{"deprecated_at": null}]}`,
		expectedStack: Stack{DeprecatedAt: time.Time{}},
	},
	"test StacksList with stacks with deprecated_at being set": {
		json:          `{"stacks": [{"deprecated_at": "2022-08-31"}]}`,
		expectedStack: Stack{DeprecatedAt: time.Date(2022, 8, 31, 0, 0, 0, 0, time.UTC)},
	},
}

func TestStackIsDeprecated(t *testing.T) {
	for message, test := range isDeprecatedCases {
		t.Run(message, func(t *testing.T) {
			stack := Stack{DeprecatedAt: test.date}
			if stack.IsDeprecated() != test.deprecated {
				t.Errorf("IsDeprecated expected to be %t, got %t", test.deprecated, stack.IsDeprecated())
			}
		})
	}
}

func TestStacksList(t *testing.T) {
	ctx := context.Background()

	for message, test := range stacksListCases {
		t.Run(message, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client, err := New(ctx, ClientConfig{})
			require.NoError(t, err)
			apiMock := httpmock.NewMockClient(ctrl)
			client.apiClient = apiMock

			apiMock.EXPECT().DoRequest(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(_ context.Context, _, res interface{}) {
				err := json.Unmarshal([]byte(test.json), &res)
				require.NoError(t, err)
			}).Return(nil)

			stacks, _ := client.StacksList(ctx)

			if stacks[0].DeprecatedAt != test.expectedStack.DeprecatedAt {
				t.Errorf("deprecation date: expected %s, got %s", test.expectedStack.DeprecatedAt.String(), stacks[0].DeprecatedAt.String())
			}
		})
	}
}
