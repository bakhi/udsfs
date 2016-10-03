package plugin

import (
	"../../udsfs"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("lorems", bql.SourceCreatorFunc(udsfs.CreateLoremSource))
	udf.MustRegisterGlobalUDSCreator("word_splitter", udf.MustConvertToUDSFCreator(udsfs.CreateWordSplitter))

}
