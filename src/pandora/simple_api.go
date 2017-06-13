package pandora

import (
	. "github.com/qiniu/pandora-go-sdk/base"
	"github.com/qiniu/pandora-go-sdk/base/config"
	"github.com/qiniu/pandora-go-sdk/logdb"
	"github.com/qiniu/pandora-go-sdk/pipeline"
	"github.com/qiniu/pandora-go-sdk/tsdb"
	"log"
)

var (
	defaultRegion = "nb"

	ak string
	sk string

	logger Logger

	defaultContainer = &pipeline.Container{
		Type:  "M16C4",
		Count: 1,
	}

	l_endpoint = "https://logdb.qiniu.com"
	p_endpoint = "https://pipeline.qiniu.com"
	t_endpoint = "https://tsdb.qiniu.com"

	p_client pipeline.PipelineAPI
	t_client tsdb.TsdbAPI
	l_client logdb.LogdbAPI
)

func Init(_ak, _sk string) {
	var err error

	ak = _ak
	sk = _sk

	logger = NewDefaultLogger()
	t_cfg := config.NewConfig().
		WithEndpoint(t_endpoint).
		WithAccessKeySecretKey(ak, sk).
		WithLogger(logger).
		WithLoggerLevel(LogDebug)

	p_cfg := config.NewConfig().
		WithEndpoint(p_endpoint).
		WithAccessKeySecretKey(ak, sk).
		WithLogger(logger).
		WithLoggerLevel(LogDebug)

	l_cfg := config.NewConfig().
		WithEndpoint(l_endpoint).
		WithAccessKeySecretKey(ak, sk).
		WithLogger(logger).
		WithLoggerLevel(LogDebug)

	p_client, err = pipeline.New(p_cfg)
	if err != nil {
		logger.Error("new pipeline client failed, err: %v", err)
		return
	}

	t_client, err = tsdb.New(t_cfg)
	if err != nil {
		logger.Error("new tsdb client failed, err: %v", err)
		return
	}

	l_client, err = logdb.New(l_cfg)
	if err != nil {
		logger.Error("new logdb client failed, err: %v", err)
		return
	}
}

func CreateLogdbExport(repo string, logdb string, doc map[string]interface{}) {
	export := pipeline.CreateExportInput{
		RepoName:   repo,
		ExportName: "log_export",
		Spec: &pipeline.ExportLogDBSpec{
			DestRepoName: logdb,
			Doc:          doc,
		},
		Whence: "newest",
	}

	err := p_client.CreateExport(&export)
	if err != nil {
		logger.Errorf("export: %s create failed, err: %v", export.ExportName, err)
		return
	}
	log.Println("create logdb export done")
}

func CreateKodoExport(repo string, bucket string, fields map[string]string, account string) {
	export := pipeline.CreateExportInput{
		RepoName:   repo,
		ExportName: "kodo_export",
		Spec: &pipeline.ExportKodoSpec{
			Bucket:         bucket,
			KeyPrefix:      "export_",
			Fields:         fields,
			Email:          account,
			AccessKey:      ak,
			Format:         "parquet",
			RotateInterval: 600,
		},
		Whence: "newest",
	}

	err := p_client.CreateExport(&export)
	if err != nil {
		logger.Errorf("export: %s create failed, err: %v", export.ExportName, err)
		return
	}
	log.Println("create kodo export done")
}

func CreateKodoExportWithAK(repo string, bucket string, fields map[string]string, account string, ak2 string) {
	export := pipeline.CreateExportInput{
		RepoName:   repo,
		ExportName: "kodo_export",
		Spec: &pipeline.ExportKodoSpec{
			Bucket:         bucket,
			KeyPrefix:      "export_",
			Fields:         fields,
			Email:          account,
			AccessKey:      ak2,
			Format:         "parquet",
			RotateInterval: 600,
		},
		Whence: "newest",
	}

	err := p_client.CreateExport(&export)
	if err != nil {
		logger.Errorf("export: %s create failed, err: %v", export.ExportName, err)
	}
}

func CreateTSRepo(repoName string) {
	createInput := &tsdb.CreateRepoInput{
		RepoName: repoName,
		Region:   defaultRegion,
	}

	err := t_client.CreateRepo(createInput)
	if err != nil {
		logger.Error("create ts repo failed", err)
		return
	}

	log.Println("create ts repo done")
}

func DeleteTSRepo(repoName string) {
	deleteRepoInput := &tsdb.DeleteRepoInput{
		RepoName: repoName,
	}

	err := t_client.DeleteRepo(deleteRepoInput)
	if err != nil {
		logger.Error("delete ts repo failed", err)
		return
	}

	log.Println("delete ts repo done")
}

func ListTSRepo() {
	listReposInput := &tsdb.ListReposInput{}

	listReposOutput, err := t_client.ListRepos(listReposInput)
	if err != nil {
		logger.Error("list ts repo failed", err)
		return
	}

	log.Println("list ts repo done")

	logger.Info(listReposOutput)
}

func CreateSeries(repoName string, seriesName string, retention string) {
	createInput := &tsdb.CreateSeriesInput{
		RepoName:   repoName,
		SeriesName: seriesName,
		Retention:  retention,
	}

	err := t_client.CreateSeries(createInput)
	if err != nil {
		logger.Error("create ts ser failed", err)
		return
	}

	log.Println("create ts ser done")
}

func DeleteSeries(repoName string, seriesName string) {
	deleteInput := &tsdb.DeleteSeriesInput{
		RepoName:   repoName,
		SeriesName: seriesName,
	}

	err := t_client.DeleteSeries(deleteInput)
	if err != nil {
		logger.Error("delete series failed", err)
		return
	}

	log.Println("delete series done")
}

func ListSeries(repoName string) {
	listSeriesInput := &tsdb.ListSeriesInput{
		RepoName: repoName,
	}

	listSeriesOutput, err := t_client.ListSeries(listSeriesInput)
	if err != nil {
		logger.Error("list ts series failed", err)
		return
	}

	log.Println("list ts series done")

	logger.Info(listSeriesOutput)
}

func CreatePipelineRepo(repoName string, schema []pipeline.RepoSchemaEntry, groupName string) {
	createRepoInput := &pipeline.CreateRepoInput{
		RepoName:  repoName,
		Region:    defaultRegion,
		Schema:    schema,
		GroupName: groupName,
	}

	err := p_client.CreateRepo(createRepoInput)
	if err != nil {
		logger.Error("create pipeline repo failed", err)
		return
	}

	log.Println("create pipeline repo done")
}

func DeletePipelineRepo(repoName string) {
	deleteInput := &pipeline.DeleteRepoInput{
		RepoName: repoName,
	}

	err := p_client.DeleteRepo(deleteInput)
	if err != nil {
		logger.Error("delete pipeline repo failed", err)
		return
	}

	log.Println("delete pipeline repo done")
}

func ListPipelineRepos() {
	listReposInput := &pipeline.ListReposInput{}

	listReposOutput, err := p_client.ListRepos(listReposInput)
	if err != nil {
		logger.Error("list pipeline repos failed", err)
		return
	}

	log.Println("list pipeline repos done")

	logger.Info(listReposOutput)
}

func IsPipelineRepoExisted(repoName string) bool {
	getOutput, err := p_client.GetRepo(&pipeline.GetRepoInput{RepoName: repoName})
	if err != nil {
		logger.Error(err)
		return false
	}

	if getOutput != nil {
		return true
	}

	return false
}

func IsTsRepoExisted(repoName string) bool {
	getOutput, err := t_client.GetRepo(&tsdb.GetRepoInput{RepoName: repoName})
	if err != nil {
		logger.Error(err)
		return false
	}

	if getOutput != nil {
		return true
	}

	return false
}

func IsTsExportExisted(repoName, tsRepoName string) bool {
	getOutput, err := p_client.GetExport(&pipeline.GetExportInput{RepoName: repoName, ExportName: repoName + "_tsexport_" + tsRepoName})
	if err != nil {
		logger.Error(err)
		return false
	}

	if getOutput != nil {
		return true
	}

	return false
}

func IsTransformExisted(repo, targetRepo string) bool {
	getOutput, err := p_client.GetTransform(&pipeline.GetTransformInput{RepoName: repo, TransformName: "transform_" + targetRepo})
	if err != nil {
		logger.Error(err)
		return false
	}

	if getOutput != nil {
		return true
	}

	return false
}

func CreateTsExport(pipelineRepoName string, tsRepoName string, seriesName string,
	tags map[string]string, fields map[string]string, timeStamp string, whence string) {
	// 在Input里面不需要指定Type，sdk会自动填充
	export := pipeline.CreateExportInput{
		RepoName:   pipelineRepoName,
		ExportName: pipelineRepoName + "_tsexport_" + tsRepoName,
		Spec: &pipeline.ExportTsdbSpec{
			DestRepoName: tsRepoName,
			SeriesName:   seriesName,
			Tags:         tags,
			Fields:       fields,
			Timestamp:    timeStamp,
		},
		Whence: whence,
	}

	err := p_client.CreateExport(&export)
	if err != nil {
		logger.Errorf("export: %s create failed, err: %v", export.ExportName, err)
	} else {
		log.Println("createExport success")
	}
}

func DeleteExport(pipelineRepoName string, exportName string) {
	exportInput := pipeline.DeleteExportInput{
		RepoName:   pipelineRepoName,
		ExportName: exportName,
	}

	err := p_client.DeleteExport(&exportInput)

	if err != nil {
		logger.Errorf("export: %s delete failed, err: %v", exportInput.ExportName, err)
	} else {
		log.Println("deleteExport success")
	}
}

func ListExports(repoName string) {
	listExportsInput := &pipeline.ListExportsInput{
		RepoName: repoName,
	}

	listExportsOutput, err := p_client.ListExports(listExportsInput)
	if err != nil {
		logger.Error("list exports failed", err)
		return
	}

	log.Println("list exports done")

	logger.Info(listExportsOutput)
}

func TsQuery() {
	q := &tsdb.QueryInput{
		RepoName: "pili2",
		Sql:      "select * from livecount2 where time > now() - 2d",
		// Sql: "select * from live where time > now() - 2h",
	}

	output, err := t_client.QueryPoints(q)
	if err != nil {
		logger.Error(output)
	}
	logger.Info(output)
}

func CreateTransformSql(repo, targetRepo, sql, interval string) {
	spec := &pipeline.TransformSpec{
		Mode:      "sql",
		Code:      sql,
		Interval:  interval,
		Container: defaultContainer,
	}
	createTransInput := &pipeline.CreateTransformInput{
		SrcRepoName:   repo,
		DestRepoName:  targetRepo,
		TransformName: "transform_" + targetRepo,
		Spec:          spec,
	}
	err := p_client.CreateTransform(createTransInput)
	if err != nil {
		logger.Error(err)
	} else {
		log.Println("create transform done")
	}
}

func UpdateTransformSql(repo, targetRepo, sql, interval string) {
	spec := &pipeline.TransformSpec{
		Mode:      "sql",
		Code:      sql,
		Interval:  interval,
		Container: defaultContainer,
	}

	updateTransInput := &pipeline.UpdateTransformInput{
		SrcRepoName:   repo,
		TransformName: "transform_" + targetRepo,
		Spec:          spec,
	}
	err := p_client.UpdateTransform(updateTransInput)
	if err != nil {
		logger.Error(err)
		return
	}
}

func CreateTransformSqlWithPlugin(repo, targetRepo, sql, interval, pluginName string, schema []pipeline.TransformPluginOutputEntry) {
	plugin := &pipeline.TransformPlugin{
		Name:   pluginName,
		Output: schema,
	}
	spec := &pipeline.TransformSpec{
		Mode:      "sql",
		Code:      sql,
		Interval:  interval,
		Container: defaultContainer,
		Plugin:    plugin,
	}
	createTransInput := &pipeline.CreateTransformInput{
		SrcRepoName:   repo,
		DestRepoName:  targetRepo,
		TransformName: "transform_" + targetRepo,
		Spec:          spec,
	}
	err := p_client.CreateTransform(createTransInput)
	if err != nil {
		logger.Error(err)
		return
	}
}

func UpdateTransformSqlWithPlugin(repo, targetRepo, sql, interval, pluginName string, schema []pipeline.TransformPluginOutputEntry) {
	plugin := &pipeline.TransformPlugin{
		Name:   pluginName,
		Output: schema,
	}
	spec := &pipeline.TransformSpec{
		Mode:      "sql",
		Code:      sql,
		Interval:  interval,
		Container: defaultContainer,
		Plugin:    plugin,
	}
	updateTransInput := &pipeline.UpdateTransformInput{
		SrcRepoName:   repo,
		TransformName: "transform_" + targetRepo,
		Spec:          spec,
	}
	err := p_client.UpdateTransform(updateTransInput)
	if err != nil {
		logger.Error(err)
		return
	}
}

func DeletePlugin(name string) {
	if err := p_client.DeletePlugin(&pipeline.DeletePluginInput{PluginName: name}); err != nil {
		logger.Error(err)
		return
	}
}

func UploadPluginFromFile(path string, pluginName string) {
	filePluginInput := &pipeline.UploadPluginFromFileInput{
		PluginName: pluginName,
		FilePath:   path,
	}

	// 从本地文件中上传plugin
	if err := p_client.UploadPluginFromFile(filePluginInput); err != nil {
		logger.Error(err)
		return
	}
}

func ListTransform(repo string) {
	listTransOutput, err := p_client.ListTransforms(&pipeline.ListTransformsInput{RepoName: repo})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(listTransOutput)
}

func DeleteTransform(repo, transform string) {
	err := p_client.DeleteTransform(&pipeline.DeleteTransformInput{RepoName: repo, TransformName: transform})
	if err != nil {
		logger.Error(err)
	}
}

func CreateLogdb(name string, schema []logdb.RepoSchemaEntry) {
	createInput := &logdb.CreateRepoInput{
		RepoName:  name,
		Region:    defaultRegion,
		Schema:    schema,
		Retention: "30d",
	}
	err := l_client.CreateRepo(createInput)
	if err != nil {
		logger.Error("create log db failed", err)
		return
	}

	log.Println("create log db done")
}
