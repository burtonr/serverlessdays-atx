package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("ServerlessDays ATX. You said: %s", string(req))
}
