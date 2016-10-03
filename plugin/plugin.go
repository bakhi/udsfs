package plugin

import (
	"github.com/udsfs"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("lorems", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	udf.MustRegisterGlobalUDSFCreator("word_splitter", udf.MustConvertToUDSFCreator(udsfs.CreateWordSplitter))
	udf.MustRegisterGlobalUDSFCreator("ticker", udf.MustConvertToUDSFCreator(udsfs.CreateTicker))
}
