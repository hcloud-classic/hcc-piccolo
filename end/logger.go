package end

import "hcc/piccolo/lib/logger"

func loggerEnd() {
	_ = logger.FpLog.Close()
}
