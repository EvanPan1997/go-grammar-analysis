// Code generated from docs/UniversalParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // UniversalParser
import "github.com/antlr4-go/antlr/v4"

// BaseUniversalParserListener is a complete listener for a parse tree produced by UniversalParser.
type BaseUniversalParserListener struct{}

var _ UniversalParserListener = &BaseUniversalParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUniversalParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUniversalParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUniversalParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUniversalParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseUniversalParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseUniversalParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseUniversalParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseUniversalParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseUniversalParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseUniversalParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterViewTableAndView is called when production viewTableAndView is entered.
func (s *BaseUniversalParserListener) EnterViewTableAndView(ctx *ViewTableAndViewContext) {}

// ExitViewTableAndView is called when production viewTableAndView is exited.
func (s *BaseUniversalParserListener) ExitViewTableAndView(ctx *ViewTableAndViewContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseUniversalParserListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseUniversalParserListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterUpdateMaterializedViewStatement is called when production updateMaterializedViewStatement is entered.
func (s *BaseUniversalParserListener) EnterUpdateMaterializedViewStatement(ctx *UpdateMaterializedViewStatementContext) {
}

// ExitUpdateMaterializedViewStatement is called when production updateMaterializedViewStatement is exited.
func (s *BaseUniversalParserListener) ExitUpdateMaterializedViewStatement(ctx *UpdateMaterializedViewStatementContext) {
}

// EnterSetMaterializedViewLifecycleStatement is called when production setMaterializedViewLifecycleStatement is entered.
func (s *BaseUniversalParserListener) EnterSetMaterializedViewLifecycleStatement(ctx *SetMaterializedViewLifecycleStatementContext) {
}

// ExitSetMaterializedViewLifecycleStatement is called when production setMaterializedViewLifecycleStatement is exited.
func (s *BaseUniversalParserListener) ExitSetMaterializedViewLifecycleStatement(ctx *SetMaterializedViewLifecycleStatementContext) {
}

// EnterUpdateMaterializedViewLifecycleStatement is called when production updateMaterializedViewLifecycleStatement is entered.
func (s *BaseUniversalParserListener) EnterUpdateMaterializedViewLifecycleStatement(ctx *UpdateMaterializedViewLifecycleStatementContext) {
}

// ExitUpdateMaterializedViewLifecycleStatement is called when production updateMaterializedViewLifecycleStatement is exited.
func (s *BaseUniversalParserListener) ExitUpdateMaterializedViewLifecycleStatement(ctx *UpdateMaterializedViewLifecycleStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseUniversalParserListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseUniversalParserListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterDropMaterializedViewPartitionStatement is called when production dropMaterializedViewPartitionStatement is entered.
func (s *BaseUniversalParserListener) EnterDropMaterializedViewPartitionStatement(ctx *DropMaterializedViewPartitionStatementContext) {
}

// ExitDropMaterializedViewPartitionStatement is called when production dropMaterializedViewPartitionStatement is exited.
func (s *BaseUniversalParserListener) ExitDropMaterializedViewPartitionStatement(ctx *DropMaterializedViewPartitionStatementContext) {
}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseUniversalParserListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseUniversalParserListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterViewColumns is called when production viewColumns is entered.
func (s *BaseUniversalParserListener) EnterViewColumns(ctx *ViewColumnsContext) {}

// ExitViewColumns is called when production viewColumns is exited.
func (s *BaseUniversalParserListener) ExitViewColumns(ctx *ViewColumnsContext) {}

// EnterRenameViewStatement is called when production renameViewStatement is entered.
func (s *BaseUniversalParserListener) EnterRenameViewStatement(ctx *RenameViewStatementContext) {}

// ExitRenameViewStatement is called when production renameViewStatement is exited.
func (s *BaseUniversalParserListener) ExitRenameViewStatement(ctx *RenameViewStatementContext) {}

// EnterChangeViewStatement is called when production changeViewStatement is entered.
func (s *BaseUniversalParserListener) EnterChangeViewStatement(ctx *ChangeViewStatementContext) {}

// ExitChangeViewStatement is called when production changeViewStatement is exited.
func (s *BaseUniversalParserListener) ExitChangeViewStatement(ctx *ChangeViewStatementContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseUniversalParserListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseUniversalParserListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseUniversalParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseUniversalParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseUniversalParserListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseUniversalParserListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterCreateTableBodyspec is called when production createTableBodyspec is entered.
func (s *BaseUniversalParserListener) EnterCreateTableBodyspec(ctx *CreateTableBodyspecContext) {}

// ExitCreateTableBodyspec is called when production createTableBodyspec is exited.
func (s *BaseUniversalParserListener) ExitCreateTableBodyspec(ctx *CreateTableBodyspecContext) {}

// EnterCreateTableItems is called when production createTableItems is entered.
func (s *BaseUniversalParserListener) EnterCreateTableItems(ctx *CreateTableItemsContext) {}

// ExitCreateTableItems is called when production createTableItems is exited.
func (s *BaseUniversalParserListener) ExitCreateTableItems(ctx *CreateTableItemsContext) {}

// EnterCreateTableItem is called when production createTableItem is entered.
func (s *BaseUniversalParserListener) EnterCreateTableItem(ctx *CreateTableItemContext) {}

// ExitCreateTableItem is called when production createTableItem is exited.
func (s *BaseUniversalParserListener) ExitCreateTableItem(ctx *CreateTableItemContext) {}

// EnterTablePartitionSpec is called when production tablePartitionSpec is entered.
func (s *BaseUniversalParserListener) EnterTablePartitionSpec(ctx *TablePartitionSpecContext) {}

// ExitTablePartitionSpec is called when production tablePartitionSpec is exited.
func (s *BaseUniversalParserListener) ExitTablePartitionSpec(ctx *TablePartitionSpecContext) {}

// EnterPartitionItems is called when production partitionItems is entered.
func (s *BaseUniversalParserListener) EnterPartitionItems(ctx *PartitionItemsContext) {}

// ExitPartitionItems is called when production partitionItems is exited.
func (s *BaseUniversalParserListener) ExitPartitionItems(ctx *PartitionItemsContext) {}

// EnterPartitionItem is called when production partitionItem is entered.
func (s *BaseUniversalParserListener) EnterPartitionItem(ctx *PartitionItemContext) {}

// ExitPartitionItem is called when production partitionItem is exited.
func (s *BaseUniversalParserListener) ExitPartitionItem(ctx *PartitionItemContext) {}

// EnterSortSpec is called when production sortSpec is entered.
func (s *BaseUniversalParserListener) EnterSortSpec(ctx *SortSpecContext) {}

// ExitSortSpec is called when production sortSpec is exited.
func (s *BaseUniversalParserListener) ExitSortSpec(ctx *SortSpecContext) {}

// EnterSkewedBySpec is called when production skewedBySpec is entered.
func (s *BaseUniversalParserListener) EnterSkewedBySpec(ctx *SkewedBySpecContext) {}

// ExitSkewedBySpec is called when production skewedBySpec is exited.
func (s *BaseUniversalParserListener) ExitSkewedBySpec(ctx *SkewedBySpecContext) {}

// EnterSkewedByColumns is called when production skewedByColumns is entered.
func (s *BaseUniversalParserListener) EnterSkewedByColumns(ctx *SkewedByColumnsContext) {}

// ExitSkewedByColumns is called when production skewedByColumns is exited.
func (s *BaseUniversalParserListener) ExitSkewedByColumns(ctx *SkewedByColumnsContext) {}

// EnterSkewedByValues is called when production skewedByValues is entered.
func (s *BaseUniversalParserListener) EnterSkewedByValues(ctx *SkewedByValuesContext) {}

// ExitSkewedByValues is called when production skewedByValues is exited.
func (s *BaseUniversalParserListener) ExitSkewedByValues(ctx *SkewedByValuesContext) {}

// EnterMultiSkewedValue is called when production multiSkewedValue is entered.
func (s *BaseUniversalParserListener) EnterMultiSkewedValue(ctx *MultiSkewedValueContext) {}

// ExitMultiSkewedValue is called when production multiSkewedValue is exited.
func (s *BaseUniversalParserListener) ExitMultiSkewedValue(ctx *MultiSkewedValueContext) {}

// EnterSimpleValue is called when production simpleValue is entered.
func (s *BaseUniversalParserListener) EnterSimpleValue(ctx *SimpleValueContext) {}

// ExitSimpleValue is called when production simpleValue is exited.
func (s *BaseUniversalParserListener) ExitSimpleValue(ctx *SimpleValueContext) {}

// EnterStorageDirectories is called when production storageDirectories is entered.
func (s *BaseUniversalParserListener) EnterStorageDirectories(ctx *StorageDirectoriesContext) {}

// ExitStorageDirectories is called when production storageDirectories is exited.
func (s *BaseUniversalParserListener) ExitStorageDirectories(ctx *StorageDirectoriesContext) {}

// EnterFormatSpec is called when production formatSpec is entered.
func (s *BaseUniversalParserListener) EnterFormatSpec(ctx *FormatSpecContext) {}

// ExitFormatSpec is called when production formatSpec is exited.
func (s *BaseUniversalParserListener) ExitFormatSpec(ctx *FormatSpecContext) {}

// EnterTableRowFormat is called when production tableRowFormat is entered.
func (s *BaseUniversalParserListener) EnterTableRowFormat(ctx *TableRowFormatContext) {}

// ExitTableRowFormat is called when production tableRowFormat is exited.
func (s *BaseUniversalParserListener) ExitTableRowFormat(ctx *TableRowFormatContext) {}

// EnterStoredBy is called when production storedBy is entered.
func (s *BaseUniversalParserListener) EnterStoredBy(ctx *StoredByContext) {}

// ExitStoredBy is called when production storedBy is exited.
func (s *BaseUniversalParserListener) ExitStoredBy(ctx *StoredByContext) {}

// EnterSerdeProperties is called when production serdeProperties is entered.
func (s *BaseUniversalParserListener) EnterSerdeProperties(ctx *SerdePropertiesContext) {}

// ExitSerdeProperties is called when production serdeProperties is exited.
func (s *BaseUniversalParserListener) ExitSerdeProperties(ctx *SerdePropertiesContext) {}

// EnterStoredSpec is called when production storedSpec is entered.
func (s *BaseUniversalParserListener) EnterStoredSpec(ctx *StoredSpecContext) {}

// ExitStoredSpec is called when production storedSpec is exited.
func (s *BaseUniversalParserListener) ExitStoredSpec(ctx *StoredSpecContext) {}

// EnterFormatType is called when production formatType is entered.
func (s *BaseUniversalParserListener) EnterFormatType(ctx *FormatTypeContext) {}

// ExitFormatType is called when production formatType is exited.
func (s *BaseUniversalParserListener) ExitFormatType(ctx *FormatTypeContext) {}

// EnterPropertiesSpec is called when production propertiesSpec is entered.
func (s *BaseUniversalParserListener) EnterPropertiesSpec(ctx *PropertiesSpecContext) {}

// ExitPropertiesSpec is called when production propertiesSpec is exited.
func (s *BaseUniversalParserListener) ExitPropertiesSpec(ctx *PropertiesSpecContext) {}

// EnterTableSelectSpec is called when production tableSelectSpec is entered.
func (s *BaseUniversalParserListener) EnterTableSelectSpec(ctx *TableSelectSpecContext) {}

// ExitTableSelectSpec is called when production tableSelectSpec is exited.
func (s *BaseUniversalParserListener) ExitTableSelectSpec(ctx *TableSelectSpecContext) {}

// EnterLifecycleSpec is called when production lifecycleSpec is entered.
func (s *BaseUniversalParserListener) EnterLifecycleSpec(ctx *LifecycleSpecContext) {}

// ExitLifecycleSpec is called when production lifecycleSpec is exited.
func (s *BaseUniversalParserListener) ExitLifecycleSpec(ctx *LifecycleSpecContext) {}

// EnterInsertQueryStatement is called when production insertQueryStatement is entered.
func (s *BaseUniversalParserListener) EnterInsertQueryStatement(ctx *InsertQueryStatementContext) {}

// ExitInsertQueryStatement is called when production insertQueryStatement is exited.
func (s *BaseUniversalParserListener) ExitInsertQueryStatement(ctx *InsertQueryStatementContext) {}

// EnterInsertValueStatement is called when production insertValueStatement is entered.
func (s *BaseUniversalParserListener) EnterInsertValueStatement(ctx *InsertValueStatementContext) {}

// ExitInsertValueStatement is called when production insertValueStatement is exited.
func (s *BaseUniversalParserListener) ExitInsertValueStatement(ctx *InsertValueStatementContext) {}

// EnterDropTablePartitions is called when production dropTablePartitions is entered.
func (s *BaseUniversalParserListener) EnterDropTablePartitions(ctx *DropTablePartitionsContext) {}

// ExitDropTablePartitions is called when production dropTablePartitions is exited.
func (s *BaseUniversalParserListener) ExitDropTablePartitions(ctx *DropTablePartitionsContext) {}

// EnterAddTablePartitions is called when production addTablePartitions is entered.
func (s *BaseUniversalParserListener) EnterAddTablePartitions(ctx *AddTablePartitionsContext) {}

// ExitAddTablePartitions is called when production addTablePartitions is exited.
func (s *BaseUniversalParserListener) ExitAddTablePartitions(ctx *AddTablePartitionsContext) {}

// EnterRenameTablePartitions is called when production renameTablePartitions is entered.
func (s *BaseUniversalParserListener) EnterRenameTablePartitions(ctx *RenameTablePartitionsContext) {}

// ExitRenameTablePartitions is called when production renameTablePartitions is exited.
func (s *BaseUniversalParserListener) ExitRenameTablePartitions(ctx *RenameTablePartitionsContext) {}

// EnterInsertOverwriteTable is called when production insertOverwriteTable is entered.
func (s *BaseUniversalParserListener) EnterInsertOverwriteTable(ctx *InsertOverwriteTableContext) {}

// ExitInsertOverwriteTable is called when production insertOverwriteTable is exited.
func (s *BaseUniversalParserListener) ExitInsertOverwriteTable(ctx *InsertOverwriteTableContext) {}

// EnterInsertIntoTable is called when production insertIntoTable is entered.
func (s *BaseUniversalParserListener) EnterInsertIntoTable(ctx *InsertIntoTableContext) {}

// ExitInsertIntoTable is called when production insertIntoTable is exited.
func (s *BaseUniversalParserListener) ExitInsertIntoTable(ctx *InsertIntoTableContext) {}

// EnterInsertValues is called when production insertValues is entered.
func (s *BaseUniversalParserListener) EnterInsertValues(ctx *InsertValuesContext) {}

// ExitInsertValues is called when production insertValues is exited.
func (s *BaseUniversalParserListener) ExitInsertValues(ctx *InsertValuesContext) {}

// EnterInsertTableColumns is called when production insertTableColumns is entered.
func (s *BaseUniversalParserListener) EnterInsertTableColumns(ctx *InsertTableColumnsContext) {}

// ExitInsertTableColumns is called when production insertTableColumns is exited.
func (s *BaseUniversalParserListener) ExitInsertTableColumns(ctx *InsertTableColumnsContext) {}

// EnterPartitionSpec is called when production partitionSpec is entered.
func (s *BaseUniversalParserListener) EnterPartitionSpec(ctx *PartitionSpecContext) {}

// ExitPartitionSpec is called when production partitionSpec is exited.
func (s *BaseUniversalParserListener) ExitPartitionSpec(ctx *PartitionSpecContext) {}

// EnterStandardPartitionVal is called when production standardPartitionVal is entered.
func (s *BaseUniversalParserListener) EnterStandardPartitionVal(ctx *StandardPartitionValContext) {}

// ExitStandardPartitionVal is called when production standardPartitionVal is exited.
func (s *BaseUniversalParserListener) ExitStandardPartitionVal(ctx *StandardPartitionValContext) {}

// EnterDefaultPartitionVal is called when production defaultPartitionVal is entered.
func (s *BaseUniversalParserListener) EnterDefaultPartitionVal(ctx *DefaultPartitionValContext) {}

// ExitDefaultPartitionVal is called when production defaultPartitionVal is exited.
func (s *BaseUniversalParserListener) ExitDefaultPartitionVal(ctx *DefaultPartitionValContext) {}

// EnterIfOption is called when production ifOption is entered.
func (s *BaseUniversalParserListener) EnterIfOption(ctx *IfOptionContext) {}

// ExitIfOption is called when production ifOption is exited.
func (s *BaseUniversalParserListener) ExitIfOption(ctx *IfOptionContext) {}

// EnterReplaceOption is called when production replaceOption is entered.
func (s *BaseUniversalParserListener) EnterReplaceOption(ctx *ReplaceOptionContext) {}

// ExitReplaceOption is called when production replaceOption is exited.
func (s *BaseUniversalParserListener) ExitReplaceOption(ctx *ReplaceOptionContext) {}

// EnterNullOption is called when production nullOption is entered.
func (s *BaseUniversalParserListener) EnterNullOption(ctx *NullOptionContext) {}

// ExitNullOption is called when production nullOption is exited.
func (s *BaseUniversalParserListener) ExitNullOption(ctx *NullOptionContext) {}

// EnterDefaultOption is called when production defaultOption is entered.
func (s *BaseUniversalParserListener) EnterDefaultOption(ctx *DefaultOptionContext) {}

// ExitDefaultOption is called when production defaultOption is exited.
func (s *BaseUniversalParserListener) ExitDefaultOption(ctx *DefaultOptionContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseUniversalParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseUniversalParserListener) ExitQuery(ctx *QueryContext) {}

// EnterCte is called when production cte is entered.
func (s *BaseUniversalParserListener) EnterCte(ctx *CteContext) {}

// ExitCte is called when production cte is exited.
func (s *BaseUniversalParserListener) ExitCte(ctx *CteContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseUniversalParserListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseUniversalParserListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseUniversalParserListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseUniversalParserListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterQueryOrganization is called when production queryOrganization is entered.
func (s *BaseUniversalParserListener) EnterQueryOrganization(ctx *QueryOrganizationContext) {}

// ExitQueryOrganization is called when production queryOrganization is exited.
func (s *BaseUniversalParserListener) ExitQueryOrganization(ctx *QueryOrganizationContext) {}

// EnterCombineOperation is called when production combineOperation is entered.
func (s *BaseUniversalParserListener) EnterCombineOperation(ctx *CombineOperationContext) {}

// ExitCombineOperation is called when production combineOperation is exited.
func (s *BaseUniversalParserListener) ExitCombineOperation(ctx *CombineOperationContext) {}

// EnterQueryTermDefault is called when production queryTermDefault is entered.
func (s *BaseUniversalParserListener) EnterQueryTermDefault(ctx *QueryTermDefaultContext) {}

// ExitQueryTermDefault is called when production queryTermDefault is exited.
func (s *BaseUniversalParserListener) ExitQueryTermDefault(ctx *QueryTermDefaultContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseUniversalParserListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseUniversalParserListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseUniversalParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseUniversalParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseUniversalParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseUniversalParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseUniversalParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseUniversalParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseUniversalParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseUniversalParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseUniversalParserListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseUniversalParserListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterFunctionTable is called when production functionTable is entered.
func (s *BaseUniversalParserListener) EnterFunctionTable(ctx *FunctionTableContext) {}

// ExitFunctionTable is called when production functionTable is exited.
func (s *BaseUniversalParserListener) ExitFunctionTable(ctx *FunctionTableContext) {}

// EnterTemporalClause is called when production temporalClause is entered.
func (s *BaseUniversalParserListener) EnterTemporalClause(ctx *TemporalClauseContext) {}

// ExitTemporalClause is called when production temporalClause is exited.
func (s *BaseUniversalParserListener) ExitTemporalClause(ctx *TemporalClauseContext) {}

// EnterSample is called when production sample is entered.
func (s *BaseUniversalParserListener) EnterSample(ctx *SampleContext) {}

// ExitSample is called when production sample is exited.
func (s *BaseUniversalParserListener) ExitSample(ctx *SampleContext) {}

// EnterSampleByPercentile is called when production sampleByPercentile is entered.
func (s *BaseUniversalParserListener) EnterSampleByPercentile(ctx *SampleByPercentileContext) {}

// ExitSampleByPercentile is called when production sampleByPercentile is exited.
func (s *BaseUniversalParserListener) ExitSampleByPercentile(ctx *SampleByPercentileContext) {}

// EnterSampleByRows is called when production sampleByRows is entered.
func (s *BaseUniversalParserListener) EnterSampleByRows(ctx *SampleByRowsContext) {}

// ExitSampleByRows is called when production sampleByRows is exited.
func (s *BaseUniversalParserListener) ExitSampleByRows(ctx *SampleByRowsContext) {}

// EnterSampleByBucket is called when production sampleByBucket is entered.
func (s *BaseUniversalParserListener) EnterSampleByBucket(ctx *SampleByBucketContext) {}

// ExitSampleByBucket is called when production sampleByBucket is exited.
func (s *BaseUniversalParserListener) ExitSampleByBucket(ctx *SampleByBucketContext) {}

// EnterSampleByBytes is called when production sampleByBytes is entered.
func (s *BaseUniversalParserListener) EnterSampleByBytes(ctx *SampleByBytesContext) {}

// ExitSampleByBytes is called when production sampleByBytes is exited.
func (s *BaseUniversalParserListener) ExitSampleByBytes(ctx *SampleByBytesContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseUniversalParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseUniversalParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterAliasedQuery is called when production aliasedQuery is entered.
func (s *BaseUniversalParserListener) EnterAliasedQuery(ctx *AliasedQueryContext) {}

// ExitAliasedQuery is called when production aliasedQuery is exited.
func (s *BaseUniversalParserListener) ExitAliasedQuery(ctx *AliasedQueryContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseUniversalParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseUniversalParserListener) ExitRelation(ctx *RelationContext) {}

// EnterRelationExtension is called when production relationExtension is entered.
func (s *BaseUniversalParserListener) EnterRelationExtension(ctx *RelationExtensionContext) {}

// ExitRelationExtension is called when production relationExtension is exited.
func (s *BaseUniversalParserListener) ExitRelationExtension(ctx *RelationExtensionContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseUniversalParserListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseUniversalParserListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterJoinCondtion is called when production joinCondtion is entered.
func (s *BaseUniversalParserListener) EnterJoinCondtion(ctx *JoinCondtionContext) {}

// ExitJoinCondtion is called when production joinCondtion is exited.
func (s *BaseUniversalParserListener) ExitJoinCondtion(ctx *JoinCondtionContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseUniversalParserListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseUniversalParserListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseUniversalParserListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseUniversalParserListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotAggregateItem is called when production pivotAggregateItem is entered.
func (s *BaseUniversalParserListener) EnterPivotAggregateItem(ctx *PivotAggregateItemContext) {}

// ExitPivotAggregateItem is called when production pivotAggregateItem is exited.
func (s *BaseUniversalParserListener) ExitPivotAggregateItem(ctx *PivotAggregateItemContext) {}

// EnterPivotColumn is called when production pivotColumn is entered.
func (s *BaseUniversalParserListener) EnterPivotColumn(ctx *PivotColumnContext) {}

// ExitPivotColumn is called when production pivotColumn is exited.
func (s *BaseUniversalParserListener) ExitPivotColumn(ctx *PivotColumnContext) {}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseUniversalParserListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseUniversalParserListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterUnpivotClause is called when production unpivotClause is entered.
func (s *BaseUniversalParserListener) EnterUnpivotClause(ctx *UnpivotClauseContext) {}

// ExitUnpivotClause is called when production unpivotClause is exited.
func (s *BaseUniversalParserListener) ExitUnpivotClause(ctx *UnpivotClauseContext) {}

// EnterUnpivotOption is called when production unpivotOption is entered.
func (s *BaseUniversalParserListener) EnterUnpivotOption(ctx *UnpivotOptionContext) {}

// ExitUnpivotOption is called when production unpivotOption is exited.
func (s *BaseUniversalParserListener) ExitUnpivotOption(ctx *UnpivotOptionContext) {}

// EnterUnpivotSingleValueColumn is called when production unpivotSingleValueColumn is entered.
func (s *BaseUniversalParserListener) EnterUnpivotSingleValueColumn(ctx *UnpivotSingleValueColumnContext) {
}

// ExitUnpivotSingleValueColumn is called when production unpivotSingleValueColumn is exited.
func (s *BaseUniversalParserListener) ExitUnpivotSingleValueColumn(ctx *UnpivotSingleValueColumnContext) {
}

// EnterUnpivotMultiValueColumn is called when production unpivotMultiValueColumn is entered.
func (s *BaseUniversalParserListener) EnterUnpivotMultiValueColumn(ctx *UnpivotMultiValueColumnContext) {
}

// ExitUnpivotMultiValueColumn is called when production unpivotMultiValueColumn is exited.
func (s *BaseUniversalParserListener) ExitUnpivotMultiValueColumn(ctx *UnpivotMultiValueColumnContext) {
}

// EnterSingleUnpivotNameColumn is called when production singleUnpivotNameColumn is entered.
func (s *BaseUniversalParserListener) EnterSingleUnpivotNameColumn(ctx *SingleUnpivotNameColumnContext) {
}

// ExitSingleUnpivotNameColumn is called when production singleUnpivotNameColumn is exited.
func (s *BaseUniversalParserListener) ExitSingleUnpivotNameColumn(ctx *SingleUnpivotNameColumnContext) {
}

// EnterMultiUnpivotNameColumn is called when production multiUnpivotNameColumn is entered.
func (s *BaseUniversalParserListener) EnterMultiUnpivotNameColumn(ctx *MultiUnpivotNameColumnContext) {
}

// ExitMultiUnpivotNameColumn is called when production multiUnpivotNameColumn is exited.
func (s *BaseUniversalParserListener) ExitMultiUnpivotNameColumn(ctx *MultiUnpivotNameColumnContext) {
}

// EnterUnpivotColumnItem is called when production unpivotColumnItem is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumnItem(ctx *UnpivotColumnItemContext) {}

// ExitUnpivotColumnItem is called when production unpivotColumnItem is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumnItem(ctx *UnpivotColumnItemContext) {}

// EnterUnpivotMultiColumnItem is called when production unpivotMultiColumnItem is entered.
func (s *BaseUniversalParserListener) EnterUnpivotMultiColumnItem(ctx *UnpivotMultiColumnItemContext) {
}

// ExitUnpivotMultiColumnItem is called when production unpivotMultiColumnItem is exited.
func (s *BaseUniversalParserListener) ExitUnpivotMultiColumnItem(ctx *UnpivotMultiColumnItemContext) {
}

// EnterUnpivotColumnAlias is called when production unpivotColumnAlias is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumnAlias(ctx *UnpivotColumnAliasContext) {}

// ExitUnpivotColumnAlias is called when production unpivotColumnAlias is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumnAlias(ctx *UnpivotColumnAliasContext) {}

// EnterUnpivotColumnAliasOption is called when production unpivotColumnAliasOption is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumnAliasOption(ctx *UnpivotColumnAliasOptionContext) {
}

// ExitUnpivotColumnAliasOption is called when production unpivotColumnAliasOption is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumnAliasOption(ctx *UnpivotColumnAliasOptionContext) {
}

// EnterUnpivotColumnAliasValueList is called when production unpivotColumnAliasValueList is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumnAliasValueList(ctx *UnpivotColumnAliasValueListContext) {
}

// ExitUnpivotColumnAliasValueList is called when production unpivotColumnAliasValueList is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumnAliasValueList(ctx *UnpivotColumnAliasValueListContext) {
}

// EnterUnpivotColumnAliasValue is called when production unpivotColumnAliasValue is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumnAliasValue(ctx *UnpivotColumnAliasValueContext) {
}

// ExitUnpivotColumnAliasValue is called when production unpivotColumnAliasValue is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumnAliasValue(ctx *UnpivotColumnAliasValueContext) {
}

// EnterUnpivotColumn is called when production unpivotColumn is entered.
func (s *BaseUniversalParserListener) EnterUnpivotColumn(ctx *UnpivotColumnContext) {}

// ExitUnpivotColumn is called when production unpivotColumn is exited.
func (s *BaseUniversalParserListener) ExitUnpivotColumn(ctx *UnpivotColumnContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *BaseUniversalParserListener) EnterLateralView(ctx *LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *BaseUniversalParserListener) ExitLateralView(ctx *LateralViewContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseUniversalParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseUniversalParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseUniversalParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseUniversalParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterGroupExpr is called when production groupExpr is entered.
func (s *BaseUniversalParserListener) EnterGroupExpr(ctx *GroupExprContext) {}

// ExitGroupExpr is called when production groupExpr is exited.
func (s *BaseUniversalParserListener) ExitGroupExpr(ctx *GroupExprContext) {}

// EnterGroupCollection is called when production groupCollection is entered.
func (s *BaseUniversalParserListener) EnterGroupCollection(ctx *GroupCollectionContext) {}

// ExitGroupCollection is called when production groupCollection is exited.
func (s *BaseUniversalParserListener) ExitGroupCollection(ctx *GroupCollectionContext) {}

// EnterGroupWith is called when production groupWith is entered.
func (s *BaseUniversalParserListener) EnterGroupWith(ctx *GroupWithContext) {}

// ExitGroupWith is called when production groupWith is exited.
func (s *BaseUniversalParserListener) ExitGroupWith(ctx *GroupWithContext) {}

// EnterGroupWithSet is called when production groupWithSet is entered.
func (s *BaseUniversalParserListener) EnterGroupWithSet(ctx *GroupWithSetContext) {}

// ExitGroupWithSet is called when production groupWithSet is exited.
func (s *BaseUniversalParserListener) ExitGroupWithSet(ctx *GroupWithSetContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseUniversalParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseUniversalParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupingElement is called when production groupingElement is entered.
func (s *BaseUniversalParserListener) EnterGroupingElement(ctx *GroupingElementContext) {}

// ExitGroupingElement is called when production groupingElement is exited.
func (s *BaseUniversalParserListener) ExitGroupingElement(ctx *GroupingElementContext) {}

// EnterGroupCube is called when production groupCube is entered.
func (s *BaseUniversalParserListener) EnterGroupCube(ctx *GroupCubeContext) {}

// ExitGroupCube is called when production groupCube is exited.
func (s *BaseUniversalParserListener) ExitGroupCube(ctx *GroupCubeContext) {}

// EnterGroupSet is called when production groupSet is entered.
func (s *BaseUniversalParserListener) EnterGroupSet(ctx *GroupSetContext) {}

// ExitGroupSet is called when production groupSet is exited.
func (s *BaseUniversalParserListener) ExitGroupSet(ctx *GroupSetContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseUniversalParserListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseUniversalParserListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseUniversalParserListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseUniversalParserListener) ExitHint(ctx *HintContext) {}

// EnterHintStatement is called when production hintStatement is entered.
func (s *BaseUniversalParserListener) EnterHintStatement(ctx *HintStatementContext) {}

// ExitHintStatement is called when production hintStatement is exited.
func (s *BaseUniversalParserListener) ExitHintStatement(ctx *HintStatementContext) {}

// EnterNamedExpressionSeq is called when production namedExpressionSeq is entered.
func (s *BaseUniversalParserListener) EnterNamedExpressionSeq(ctx *NamedExpressionSeqContext) {}

// ExitNamedExpressionSeq is called when production namedExpressionSeq is exited.
func (s *BaseUniversalParserListener) ExitNamedExpressionSeq(ctx *NamedExpressionSeqContext) {}

// EnterQueryStars is called when production queryStars is entered.
func (s *BaseUniversalParserListener) EnterQueryStars(ctx *QueryStarsContext) {}

// ExitQueryStars is called when production queryStars is exited.
func (s *BaseUniversalParserListener) ExitQueryStars(ctx *QueryStarsContext) {}

// EnterQueryColumn is called when production queryColumn is entered.
func (s *BaseUniversalParserListener) EnterQueryColumn(ctx *QueryColumnContext) {}

// ExitQueryColumn is called when production queryColumn is exited.
func (s *BaseUniversalParserListener) ExitQueryColumn(ctx *QueryColumnContext) {}

// EnterAllQueryColumns is called when production allQueryColumns is entered.
func (s *BaseUniversalParserListener) EnterAllQueryColumns(ctx *AllQueryColumnsContext) {}

// ExitAllQueryColumns is called when production allQueryColumns is exited.
func (s *BaseUniversalParserListener) ExitAllQueryColumns(ctx *AllQueryColumnsContext) {}

// EnterExpressionSeqs is called when production expressionSeqs is entered.
func (s *BaseUniversalParserListener) EnterExpressionSeqs(ctx *ExpressionSeqsContext) {}

// ExitExpressionSeqs is called when production expressionSeqs is exited.
func (s *BaseUniversalParserListener) ExitExpressionSeqs(ctx *ExpressionSeqsContext) {}

// EnterExpressionSeq is called when production expressionSeq is entered.
func (s *BaseUniversalParserListener) EnterExpressionSeq(ctx *ExpressionSeqContext) {}

// ExitExpressionSeq is called when production expressionSeq is exited.
func (s *BaseUniversalParserListener) ExitExpressionSeq(ctx *ExpressionSeqContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUniversalParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUniversalParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseUniversalParserListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseUniversalParserListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseUniversalParserListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseUniversalParserListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseUniversalParserListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseUniversalParserListener) ExitExists(ctx *ExistsContext) {}

// EnterDefaultBinary is called when production defaultBinary is entered.
func (s *BaseUniversalParserListener) EnterDefaultBinary(ctx *DefaultBinaryContext) {}

// ExitDefaultBinary is called when production defaultBinary is exited.
func (s *BaseUniversalParserListener) ExitDefaultBinary(ctx *DefaultBinaryContext) {}

// EnterBooleanPriority is called when production booleanPriority is entered.
func (s *BaseUniversalParserListener) EnterBooleanPriority(ctx *BooleanPriorityContext) {}

// ExitBooleanPriority is called when production booleanPriority is exited.
func (s *BaseUniversalParserListener) ExitBooleanPriority(ctx *BooleanPriorityContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseUniversalParserListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseUniversalParserListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterBetweenPredicated is called when production betweenPredicated is entered.
func (s *BaseUniversalParserListener) EnterBetweenPredicated(ctx *BetweenPredicatedContext) {}

// ExitBetweenPredicated is called when production betweenPredicated is exited.
func (s *BaseUniversalParserListener) ExitBetweenPredicated(ctx *BetweenPredicatedContext) {}

// EnterInValuePredicated is called when production inValuePredicated is entered.
func (s *BaseUniversalParserListener) EnterInValuePredicated(ctx *InValuePredicatedContext) {}

// ExitInValuePredicated is called when production inValuePredicated is exited.
func (s *BaseUniversalParserListener) ExitInValuePredicated(ctx *InValuePredicatedContext) {}

// EnterInQueryPredicated is called when production inQueryPredicated is entered.
func (s *BaseUniversalParserListener) EnterInQueryPredicated(ctx *InQueryPredicatedContext) {}

// ExitInQueryPredicated is called when production inQueryPredicated is exited.
func (s *BaseUniversalParserListener) ExitInQueryPredicated(ctx *InQueryPredicatedContext) {}

// EnterRlikePredicated is called when production rlikePredicated is entered.
func (s *BaseUniversalParserListener) EnterRlikePredicated(ctx *RlikePredicatedContext) {}

// ExitRlikePredicated is called when production rlikePredicated is exited.
func (s *BaseUniversalParserListener) ExitRlikePredicated(ctx *RlikePredicatedContext) {}

// EnterRegexpPredicated is called when production regexpPredicated is entered.
func (s *BaseUniversalParserListener) EnterRegexpPredicated(ctx *RegexpPredicatedContext) {}

// ExitRegexpPredicated is called when production regexpPredicated is exited.
func (s *BaseUniversalParserListener) ExitRegexpPredicated(ctx *RegexpPredicatedContext) {}

// EnterLikeQuantifierPredicated is called when production likeQuantifierPredicated is entered.
func (s *BaseUniversalParserListener) EnterLikeQuantifierPredicated(ctx *LikeQuantifierPredicatedContext) {
}

// ExitLikeQuantifierPredicated is called when production likeQuantifierPredicated is exited.
func (s *BaseUniversalParserListener) ExitLikeQuantifierPredicated(ctx *LikeQuantifierPredicatedContext) {
}

// EnterLikePatternPredicated is called when production likePatternPredicated is entered.
func (s *BaseUniversalParserListener) EnterLikePatternPredicated(ctx *LikePatternPredicatedContext) {}

// ExitLikePatternPredicated is called when production likePatternPredicated is exited.
func (s *BaseUniversalParserListener) ExitLikePatternPredicated(ctx *LikePatternPredicatedContext) {}

// EnterNullPredicated is called when production nullPredicated is entered.
func (s *BaseUniversalParserListener) EnterNullPredicated(ctx *NullPredicatedContext) {}

// ExitNullPredicated is called when production nullPredicated is exited.
func (s *BaseUniversalParserListener) ExitNullPredicated(ctx *NullPredicatedContext) {}

// EnterBooleanPredicated is called when production booleanPredicated is entered.
func (s *BaseUniversalParserListener) EnterBooleanPredicated(ctx *BooleanPredicatedContext) {}

// ExitBooleanPredicated is called when production booleanPredicated is exited.
func (s *BaseUniversalParserListener) ExitBooleanPredicated(ctx *BooleanPredicatedContext) {}

// EnterDistinctFromPredicated is called when production distinctFromPredicated is entered.
func (s *BaseUniversalParserListener) EnterDistinctFromPredicated(ctx *DistinctFromPredicatedContext) {
}

// ExitDistinctFromPredicated is called when production distinctFromPredicated is exited.
func (s *BaseUniversalParserListener) ExitDistinctFromPredicated(ctx *DistinctFromPredicatedContext) {
}

// EnterDefaultUnary is called when production defaultUnary is entered.
func (s *BaseUniversalParserListener) EnterDefaultUnary(ctx *DefaultUnaryContext) {}

// ExitDefaultUnary is called when production defaultUnary is exited.
func (s *BaseUniversalParserListener) ExitDefaultUnary(ctx *DefaultUnaryContext) {}

// EnterValuePriority is called when production valuePriority is entered.
func (s *BaseUniversalParserListener) EnterValuePriority(ctx *ValuePriorityContext) {}

// ExitValuePriority is called when production valuePriority is exited.
func (s *BaseUniversalParserListener) ExitValuePriority(ctx *ValuePriorityContext) {}

// EnterArithmeticUnaryReverse is called when production arithmeticUnaryReverse is entered.
func (s *BaseUniversalParserListener) EnterArithmeticUnaryReverse(ctx *ArithmeticUnaryReverseContext) {
}

// ExitArithmeticUnaryReverse is called when production arithmeticUnaryReverse is exited.
func (s *BaseUniversalParserListener) ExitArithmeticUnaryReverse(ctx *ArithmeticUnaryReverseContext) {
}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseUniversalParserListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseUniversalParserListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterCompareBinary is called when production compareBinary is entered.
func (s *BaseUniversalParserListener) EnterCompareBinary(ctx *CompareBinaryContext) {}

// ExitCompareBinary is called when production compareBinary is exited.
func (s *BaseUniversalParserListener) ExitCompareBinary(ctx *CompareBinaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseUniversalParserListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseUniversalParserListener) ExitDereference(ctx *DereferenceContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseUniversalParserListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseUniversalParserListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseUniversalParserListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseUniversalParserListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterFixValue is called when production fixValue is entered.
func (s *BaseUniversalParserListener) EnterFixValue(ctx *FixValueContext) {}

// ExitFixValue is called when production fixValue is exited.
func (s *BaseUniversalParserListener) ExitFixValue(ctx *FixValueContext) {}

// EnterSubQueryExpression is called when production subQueryExpression is entered.
func (s *BaseUniversalParserListener) EnterSubQueryExpression(ctx *SubQueryExpressionContext) {}

// ExitSubQueryExpression is called when production subQueryExpression is exited.
func (s *BaseUniversalParserListener) ExitSubQueryExpression(ctx *SubQueryExpressionContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseUniversalParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseUniversalParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterSubScript is called when production subScript is entered.
func (s *BaseUniversalParserListener) EnterSubScript(ctx *SubScriptContext) {}

// ExitSubScript is called when production subScript is exited.
func (s *BaseUniversalParserListener) ExitSubScript(ctx *SubScriptContext) {}

// EnterCaseExpress is called when production caseExpress is entered.
func (s *BaseUniversalParserListener) EnterCaseExpress(ctx *CaseExpressContext) {}

// ExitCaseExpress is called when production caseExpress is exited.
func (s *BaseUniversalParserListener) ExitCaseExpress(ctx *CaseExpressContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseUniversalParserListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseUniversalParserListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseUniversalParserListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseUniversalParserListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseUniversalParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseUniversalParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterElseClause is called when production elseClause is entered.
func (s *BaseUniversalParserListener) EnterElseClause(ctx *ElseClauseContext) {}

// ExitElseClause is called when production elseClause is exited.
func (s *BaseUniversalParserListener) ExitElseClause(ctx *ElseClauseContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseUniversalParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseUniversalParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterSpecFunction is called when production specFunction is entered.
func (s *BaseUniversalParserListener) EnterSpecFunction(ctx *SpecFunctionContext) {}

// ExitSpecFunction is called when production specFunction is exited.
func (s *BaseUniversalParserListener) ExitSpecFunction(ctx *SpecFunctionContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseUniversalParserListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseUniversalParserListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterStructFunction is called when production structFunction is entered.
func (s *BaseUniversalParserListener) EnterStructFunction(ctx *StructFunctionContext) {}

// ExitStructFunction is called when production structFunction is exited.
func (s *BaseUniversalParserListener) ExitStructFunction(ctx *StructFunctionContext) {}

// EnterFirstFunction is called when production firstFunction is entered.
func (s *BaseUniversalParserListener) EnterFirstFunction(ctx *FirstFunctionContext) {}

// ExitFirstFunction is called when production firstFunction is exited.
func (s *BaseUniversalParserListener) ExitFirstFunction(ctx *FirstFunctionContext) {}

// EnterAnyValueFunction is called when production anyValueFunction is entered.
func (s *BaseUniversalParserListener) EnterAnyValueFunction(ctx *AnyValueFunctionContext) {}

// ExitAnyValueFunction is called when production anyValueFunction is exited.
func (s *BaseUniversalParserListener) ExitAnyValueFunction(ctx *AnyValueFunctionContext) {}

// EnterLastFunction is called when production lastFunction is entered.
func (s *BaseUniversalParserListener) EnterLastFunction(ctx *LastFunctionContext) {}

// ExitLastFunction is called when production lastFunction is exited.
func (s *BaseUniversalParserListener) ExitLastFunction(ctx *LastFunctionContext) {}

// EnterSubQueryExpress is called when production subQueryExpress is entered.
func (s *BaseUniversalParserListener) EnterSubQueryExpress(ctx *SubQueryExpressContext) {}

// ExitSubQueryExpress is called when production subQueryExpress is exited.
func (s *BaseUniversalParserListener) ExitSubQueryExpress(ctx *SubQueryExpressContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BaseUniversalParserListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BaseUniversalParserListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterExtractFunction is called when production extractFunction is entered.
func (s *BaseUniversalParserListener) EnterExtractFunction(ctx *ExtractFunctionContext) {}

// ExitExtractFunction is called when production extractFunction is exited.
func (s *BaseUniversalParserListener) ExitExtractFunction(ctx *ExtractFunctionContext) {}

// EnterSubstringFunction is called when production substringFunction is entered.
func (s *BaseUniversalParserListener) EnterSubstringFunction(ctx *SubstringFunctionContext) {}

// ExitSubstringFunction is called when production substringFunction is exited.
func (s *BaseUniversalParserListener) ExitSubstringFunction(ctx *SubstringFunctionContext) {}

// EnterOverlayFunction is called when production overlayFunction is entered.
func (s *BaseUniversalParserListener) EnterOverlayFunction(ctx *OverlayFunctionContext) {}

// ExitOverlayFunction is called when production overlayFunction is exited.
func (s *BaseUniversalParserListener) ExitOverlayFunction(ctx *OverlayFunctionContext) {}

// EnterPercentileFunction is called when production percentileFunction is entered.
func (s *BaseUniversalParserListener) EnterPercentileFunction(ctx *PercentileFunctionContext) {}

// ExitPercentileFunction is called when production percentileFunction is exited.
func (s *BaseUniversalParserListener) ExitPercentileFunction(ctx *PercentileFunctionContext) {}

// EnterTransArrayFunction is called when production transArrayFunction is entered.
func (s *BaseUniversalParserListener) EnterTransArrayFunction(ctx *TransArrayFunctionContext) {}

// ExitTransArrayFunction is called when production transArrayFunction is exited.
func (s *BaseUniversalParserListener) ExitTransArrayFunction(ctx *TransArrayFunctionContext) {}

// EnterStandardFunction is called when production standardFunction is entered.
func (s *BaseUniversalParserListener) EnterStandardFunction(ctx *StandardFunctionContext) {}

// ExitStandardFunction is called when production standardFunction is exited.
func (s *BaseUniversalParserListener) ExitStandardFunction(ctx *StandardFunctionContext) {}

// EnterOverSpec is called when production overSpec is entered.
func (s *BaseUniversalParserListener) EnterOverSpec(ctx *OverSpecContext) {}

// ExitOverSpec is called when production overSpec is exited.
func (s *BaseUniversalParserListener) ExitOverSpec(ctx *OverSpecContext) {}

// EnterFilterSpec is called when production filterSpec is entered.
func (s *BaseUniversalParserListener) EnterFilterSpec(ctx *FilterSpecContext) {}

// ExitFilterSpec is called when production filterSpec is exited.
func (s *BaseUniversalParserListener) ExitFilterSpec(ctx *FilterSpecContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseUniversalParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseUniversalParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseUniversalParserListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseUniversalParserListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseUniversalParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseUniversalParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterTypeLiteral is called when production typeLiteral is entered.
func (s *BaseUniversalParserListener) EnterTypeLiteral(ctx *TypeLiteralContext) {}

// ExitTypeLiteral is called when production typeLiteral is exited.
func (s *BaseUniversalParserListener) ExitTypeLiteral(ctx *TypeLiteralContext) {}

// EnterBooleaniteral is called when production booleaniteral is entered.
func (s *BaseUniversalParserListener) EnterBooleaniteral(ctx *BooleaniteralContext) {}

// ExitBooleaniteral is called when production booleaniteral is exited.
func (s *BaseUniversalParserListener) ExitBooleaniteral(ctx *BooleaniteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseUniversalParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseUniversalParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterCompareOperator is called when production compareOperator is entered.
func (s *BaseUniversalParserListener) EnterCompareOperator(ctx *CompareOperatorContext) {}

// ExitCompareOperator is called when production compareOperator is exited.
func (s *BaseUniversalParserListener) ExitCompareOperator(ctx *CompareOperatorContext) {}

// EnterArrayDataType is called when production arrayDataType is entered.
func (s *BaseUniversalParserListener) EnterArrayDataType(ctx *ArrayDataTypeContext) {}

// ExitArrayDataType is called when production arrayDataType is exited.
func (s *BaseUniversalParserListener) ExitArrayDataType(ctx *ArrayDataTypeContext) {}

// EnterMapDataType is called when production mapDataType is entered.
func (s *BaseUniversalParserListener) EnterMapDataType(ctx *MapDataTypeContext) {}

// ExitMapDataType is called when production mapDataType is exited.
func (s *BaseUniversalParserListener) ExitMapDataType(ctx *MapDataTypeContext) {}

// EnterStructDataType is called when production structDataType is entered.
func (s *BaseUniversalParserListener) EnterStructDataType(ctx *StructDataTypeContext) {}

// ExitStructDataType is called when production structDataType is exited.
func (s *BaseUniversalParserListener) ExitStructDataType(ctx *StructDataTypeContext) {}

// EnterIntervalDataType is called when production intervalDataType is entered.
func (s *BaseUniversalParserListener) EnterIntervalDataType(ctx *IntervalDataTypeContext) {}

// ExitIntervalDataType is called when production intervalDataType is exited.
func (s *BaseUniversalParserListener) ExitIntervalDataType(ctx *IntervalDataTypeContext) {}

// EnterPrimitiveDataType is called when production primitiveDataType is entered.
func (s *BaseUniversalParserListener) EnterPrimitiveDataType(ctx *PrimitiveDataTypeContext) {}

// ExitPrimitiveDataType is called when production primitiveDataType is exited.
func (s *BaseUniversalParserListener) ExitPrimitiveDataType(ctx *PrimitiveDataTypeContext) {}

// EnterComplexColTypeList is called when production complexColTypeList is entered.
func (s *BaseUniversalParserListener) EnterComplexColTypeList(ctx *ComplexColTypeListContext) {}

// ExitComplexColTypeList is called when production complexColTypeList is exited.
func (s *BaseUniversalParserListener) ExitComplexColTypeList(ctx *ComplexColTypeListContext) {}

// EnterComplexColType is called when production complexColType is entered.
func (s *BaseUniversalParserListener) EnterComplexColType(ctx *ComplexColTypeContext) {}

// ExitComplexColType is called when production complexColType is exited.
func (s *BaseUniversalParserListener) ExitComplexColType(ctx *ComplexColTypeContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseUniversalParserListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseUniversalParserListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseUniversalParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseUniversalParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterMultiUnitsInterval is called when production multiUnitsInterval is entered.
func (s *BaseUniversalParserListener) EnterMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// ExitMultiUnitsInterval is called when production multiUnitsInterval is exited.
func (s *BaseUniversalParserListener) ExitMultiUnitsInterval(ctx *MultiUnitsIntervalContext) {}

// EnterUnitToUnitInterval is called when production unitToUnitInterval is entered.
func (s *BaseUniversalParserListener) EnterUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// ExitUnitToUnitInterval is called when production unitToUnitInterval is exited.
func (s *BaseUniversalParserListener) ExitUnitToUnitInterval(ctx *UnitToUnitIntervalContext) {}

// EnterMultiUnits is called when production multiUnits is entered.
func (s *BaseUniversalParserListener) EnterMultiUnits(ctx *MultiUnitsContext) {}

// ExitMultiUnits is called when production multiUnits is exited.
func (s *BaseUniversalParserListener) ExitMultiUnits(ctx *MultiUnitsContext) {}

// EnterUnitToUnit is called when production unitToUnit is entered.
func (s *BaseUniversalParserListener) EnterUnitToUnit(ctx *UnitToUnitContext) {}

// ExitUnitToUnit is called when production unitToUnit is exited.
func (s *BaseUniversalParserListener) ExitUnitToUnit(ctx *UnitToUnitContext) {}

// EnterCommentSpec is called when production commentSpec is entered.
func (s *BaseUniversalParserListener) EnterCommentSpec(ctx *CommentSpecContext) {}

// ExitCommentSpec is called when production commentSpec is exited.
func (s *BaseUniversalParserListener) ExitCommentSpec(ctx *CommentSpecContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseUniversalParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseUniversalParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterNamedWindow is called when production namedWindow is entered.
func (s *BaseUniversalParserListener) EnterNamedWindow(ctx *NamedWindowContext) {}

// ExitNamedWindow is called when production namedWindow is exited.
func (s *BaseUniversalParserListener) ExitNamedWindow(ctx *NamedWindowContext) {}

// EnterWindowRef is called when production windowRef is entered.
func (s *BaseUniversalParserListener) EnterWindowRef(ctx *WindowRefContext) {}

// ExitWindowRef is called when production windowRef is exited.
func (s *BaseUniversalParserListener) ExitWindowRef(ctx *WindowRefContext) {}

// EnterWindowDef is called when production windowDef is entered.
func (s *BaseUniversalParserListener) EnterWindowDef(ctx *WindowDefContext) {}

// ExitWindowDef is called when production windowDef is exited.
func (s *BaseUniversalParserListener) ExitWindowDef(ctx *WindowDefContext) {}

// EnterStandardWindowFrame is called when production standardWindowFrame is entered.
func (s *BaseUniversalParserListener) EnterStandardWindowFrame(ctx *StandardWindowFrameContext) {}

// ExitStandardWindowFrame is called when production standardWindowFrame is exited.
func (s *BaseUniversalParserListener) ExitStandardWindowFrame(ctx *StandardWindowFrameContext) {}

// EnterBetweenWindowFrame is called when production betweenWindowFrame is entered.
func (s *BaseUniversalParserListener) EnterBetweenWindowFrame(ctx *BetweenWindowFrameContext) {}

// ExitBetweenWindowFrame is called when production betweenWindowFrame is exited.
func (s *BaseUniversalParserListener) ExitBetweenWindowFrame(ctx *BetweenWindowFrameContext) {}

// EnterUnboundedPreceding is called when production unboundedPreceding is entered.
func (s *BaseUniversalParserListener) EnterUnboundedPreceding(ctx *UnboundedPrecedingContext) {}

// ExitUnboundedPreceding is called when production unboundedPreceding is exited.
func (s *BaseUniversalParserListener) ExitUnboundedPreceding(ctx *UnboundedPrecedingContext) {}

// EnterCurrentRow is called when production currentRow is entered.
func (s *BaseUniversalParserListener) EnterCurrentRow(ctx *CurrentRowContext) {}

// ExitCurrentRow is called when production currentRow is exited.
func (s *BaseUniversalParserListener) ExitCurrentRow(ctx *CurrentRowContext) {}

// EnterPreceding is called when production preceding is entered.
func (s *BaseUniversalParserListener) EnterPreceding(ctx *PrecedingContext) {}

// ExitPreceding is called when production preceding is exited.
func (s *BaseUniversalParserListener) ExitPreceding(ctx *PrecedingContext) {}

// EnterTableProvider is called when production tableProvider is entered.
func (s *BaseUniversalParserListener) EnterTableProvider(ctx *TableProviderContext) {}

// ExitTableProvider is called when production tableProvider is exited.
func (s *BaseUniversalParserListener) ExitTableProvider(ctx *TableProviderContext) {}

// EnterCreateFileFormat is called when production createFileFormat is entered.
func (s *BaseUniversalParserListener) EnterCreateFileFormat(ctx *CreateFileFormatContext) {}

// ExitCreateFileFormat is called when production createFileFormat is exited.
func (s *BaseUniversalParserListener) ExitCreateFileFormat(ctx *CreateFileFormatContext) {}

// EnterTableFileFormat is called when production tableFileFormat is entered.
func (s *BaseUniversalParserListener) EnterTableFileFormat(ctx *TableFileFormatContext) {}

// ExitTableFileFormat is called when production tableFileFormat is exited.
func (s *BaseUniversalParserListener) ExitTableFileFormat(ctx *TableFileFormatContext) {}

// EnterGenericFileFormat is called when production genericFileFormat is entered.
func (s *BaseUniversalParserListener) EnterGenericFileFormat(ctx *GenericFileFormatContext) {}

// ExitGenericFileFormat is called when production genericFileFormat is exited.
func (s *BaseUniversalParserListener) ExitGenericFileFormat(ctx *GenericFileFormatContext) {}

// EnterStorageHandler is called when production storageHandler is entered.
func (s *BaseUniversalParserListener) EnterStorageHandler(ctx *StorageHandlerContext) {}

// ExitStorageHandler is called when production storageHandler is exited.
func (s *BaseUniversalParserListener) ExitStorageHandler(ctx *StorageHandlerContext) {}

// EnterRowFormatSerde is called when production rowFormatSerde is entered.
func (s *BaseUniversalParserListener) EnterRowFormatSerde(ctx *RowFormatSerdeContext) {}

// ExitRowFormatSerde is called when production rowFormatSerde is exited.
func (s *BaseUniversalParserListener) ExitRowFormatSerde(ctx *RowFormatSerdeContext) {}

// EnterRowFormatDelimited is called when production rowFormatDelimited is entered.
func (s *BaseUniversalParserListener) EnterRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// ExitRowFormatDelimited is called when production rowFormatDelimited is exited.
func (s *BaseUniversalParserListener) ExitRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseUniversalParserListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseUniversalParserListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseUniversalParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseUniversalParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterPropertyKey is called when production propertyKey is entered.
func (s *BaseUniversalParserListener) EnterPropertyKey(ctx *PropertyKeyContext) {}

// ExitPropertyKey is called when production propertyKey is exited.
func (s *BaseUniversalParserListener) ExitPropertyKey(ctx *PropertyKeyContext) {}

// EnterPropertyValue is called when production propertyValue is entered.
func (s *BaseUniversalParserListener) EnterPropertyValue(ctx *PropertyValueContext) {}

// ExitPropertyValue is called when production propertyValue is exited.
func (s *BaseUniversalParserListener) ExitPropertyValue(ctx *PropertyValueContext) {}

// EnterStandardArgs is called when production standardArgs is entered.
func (s *BaseUniversalParserListener) EnterStandardArgs(ctx *StandardArgsContext) {}

// ExitStandardArgs is called when production standardArgs is exited.
func (s *BaseUniversalParserListener) ExitStandardArgs(ctx *StandardArgsContext) {}

// EnterStars is called when production stars is entered.
func (s *BaseUniversalParserListener) EnterStars(ctx *StarsContext) {}

// ExitStars is called when production stars is exited.
func (s *BaseUniversalParserListener) ExitStars(ctx *StarsContext) {}

// EnterOrderBy is called when production orderBy is entered.
func (s *BaseUniversalParserListener) EnterOrderBy(ctx *OrderByContext) {}

// ExitOrderBy is called when production orderBy is exited.
func (s *BaseUniversalParserListener) ExitOrderBy(ctx *OrderByContext) {}

// EnterSortBy is called when production sortBy is entered.
func (s *BaseUniversalParserListener) EnterSortBy(ctx *SortByContext) {}

// ExitSortBy is called when production sortBy is exited.
func (s *BaseUniversalParserListener) ExitSortBy(ctx *SortByContext) {}

// EnterDistributeBy is called when production distributeBy is entered.
func (s *BaseUniversalParserListener) EnterDistributeBy(ctx *DistributeByContext) {}

// ExitDistributeBy is called when production distributeBy is exited.
func (s *BaseUniversalParserListener) ExitDistributeBy(ctx *DistributeByContext) {}

// EnterPartitionBy is called when production partitionBy is entered.
func (s *BaseUniversalParserListener) EnterPartitionBy(ctx *PartitionByContext) {}

// ExitPartitionBy is called when production partitionBy is exited.
func (s *BaseUniversalParserListener) ExitPartitionBy(ctx *PartitionByContext) {}

// EnterClusterBy is called when production clusterBy is entered.
func (s *BaseUniversalParserListener) EnterClusterBy(ctx *ClusterByContext) {}

// ExitClusterBy is called when production clusterBy is exited.
func (s *BaseUniversalParserListener) ExitClusterBy(ctx *ClusterByContext) {}

// EnterBucketsOption is called when production bucketsOption is entered.
func (s *BaseUniversalParserListener) EnterBucketsOption(ctx *BucketsOptionContext) {}

// ExitBucketsOption is called when production bucketsOption is exited.
func (s *BaseUniversalParserListener) ExitBucketsOption(ctx *BucketsOptionContext) {}

// EnterLimitItem is called when production limitItem is entered.
func (s *BaseUniversalParserListener) EnterLimitItem(ctx *LimitItemContext) {}

// ExitLimitItem is called when production limitItem is exited.
func (s *BaseUniversalParserListener) ExitLimitItem(ctx *LimitItemContext) {}

// EnterOffsetItem is called when production offsetItem is entered.
func (s *BaseUniversalParserListener) EnterOffsetItem(ctx *OffsetItemContext) {}

// ExitOffsetItem is called when production offsetItem is exited.
func (s *BaseUniversalParserListener) ExitOffsetItem(ctx *OffsetItemContext) {}

// EnterSortItems is called when production sortItems is entered.
func (s *BaseUniversalParserListener) EnterSortItems(ctx *SortItemsContext) {}

// ExitSortItems is called when production sortItems is exited.
func (s *BaseUniversalParserListener) ExitSortItems(ctx *SortItemsContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseUniversalParserListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseUniversalParserListener) ExitSortItem(ctx *SortItemContext) {}

// EnterSortType is called when production sortType is entered.
func (s *BaseUniversalParserListener) EnterSortType(ctx *SortTypeContext) {}

// ExitSortType is called when production sortType is exited.
func (s *BaseUniversalParserListener) ExitSortType(ctx *SortTypeContext) {}

// EnterAliasSpec is called when production aliasSpec is entered.
func (s *BaseUniversalParserListener) EnterAliasSpec(ctx *AliasSpecContext) {}

// ExitAliasSpec is called when production aliasSpec is exited.
func (s *BaseUniversalParserListener) ExitAliasSpec(ctx *AliasSpecContext) {}

// EnterLocationSpec is called when production locationSpec is entered.
func (s *BaseUniversalParserListener) EnterLocationSpec(ctx *LocationSpecContext) {}

// ExitLocationSpec is called when production locationSpec is exited.
func (s *BaseUniversalParserListener) ExitLocationSpec(ctx *LocationSpecContext) {}

// EnterMultiIdentifierSeq is called when production multiIdentifierSeq is entered.
func (s *BaseUniversalParserListener) EnterMultiIdentifierSeq(ctx *MultiIdentifierSeqContext) {}

// ExitMultiIdentifierSeq is called when production multiIdentifierSeq is exited.
func (s *BaseUniversalParserListener) ExitMultiIdentifierSeq(ctx *MultiIdentifierSeqContext) {}

// EnterMultiIdentifier is called when production multiIdentifier is entered.
func (s *BaseUniversalParserListener) EnterMultiIdentifier(ctx *MultiIdentifierContext) {}

// ExitMultiIdentifier is called when production multiIdentifier is exited.
func (s *BaseUniversalParserListener) ExitMultiIdentifier(ctx *MultiIdentifierContext) {}

// EnterIdentifierSeq is called when production identifierSeq is entered.
func (s *BaseUniversalParserListener) EnterIdentifierSeq(ctx *IdentifierSeqContext) {}

// ExitIdentifierSeq is called when production identifierSeq is exited.
func (s *BaseUniversalParserListener) ExitIdentifierSeq(ctx *IdentifierSeqContext) {}

// EnterUnquotedIdentifierAlternative is called when production unquotedIdentifierAlternative is entered.
func (s *BaseUniversalParserListener) EnterUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) {
}

// ExitUnquotedIdentifierAlternative is called when production unquotedIdentifierAlternative is exited.
func (s *BaseUniversalParserListener) ExitUnquotedIdentifierAlternative(ctx *UnquotedIdentifierAlternativeContext) {
}

// EnterQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is entered.
func (s *BaseUniversalParserListener) EnterQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// ExitQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is exited.
func (s *BaseUniversalParserListener) ExitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseUniversalParserListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseUniversalParserListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterExponentLiteral is called when production exponentLiteral is entered.
func (s *BaseUniversalParserListener) EnterExponentLiteral(ctx *ExponentLiteralContext) {}

// ExitExponentLiteral is called when production exponentLiteral is exited.
func (s *BaseUniversalParserListener) ExitExponentLiteral(ctx *ExponentLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseUniversalParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseUniversalParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterLegacyDecimalLiteral is called when production legacyDecimalLiteral is entered.
func (s *BaseUniversalParserListener) EnterLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) {}

// ExitLegacyDecimalLiteral is called when production legacyDecimalLiteral is exited.
func (s *BaseUniversalParserListener) ExitLegacyDecimalLiteral(ctx *LegacyDecimalLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseUniversalParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseUniversalParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterBigIntLiteral is called when production bigIntLiteral is entered.
func (s *BaseUniversalParserListener) EnterBigIntLiteral(ctx *BigIntLiteralContext) {}

// ExitBigIntLiteral is called when production bigIntLiteral is exited.
func (s *BaseUniversalParserListener) ExitBigIntLiteral(ctx *BigIntLiteralContext) {}

// EnterSmallIntLiteral is called when production smallIntLiteral is entered.
func (s *BaseUniversalParserListener) EnterSmallIntLiteral(ctx *SmallIntLiteralContext) {}

// ExitSmallIntLiteral is called when production smallIntLiteral is exited.
func (s *BaseUniversalParserListener) ExitSmallIntLiteral(ctx *SmallIntLiteralContext) {}

// EnterTinyIntLiteral is called when production tinyIntLiteral is entered.
func (s *BaseUniversalParserListener) EnterTinyIntLiteral(ctx *TinyIntLiteralContext) {}

// ExitTinyIntLiteral is called when production tinyIntLiteral is exited.
func (s *BaseUniversalParserListener) ExitTinyIntLiteral(ctx *TinyIntLiteralContext) {}

// EnterDoubleLiteral is called when production doubleLiteral is entered.
func (s *BaseUniversalParserListener) EnterDoubleLiteral(ctx *DoubleLiteralContext) {}

// ExitDoubleLiteral is called when production doubleLiteral is exited.
func (s *BaseUniversalParserListener) ExitDoubleLiteral(ctx *DoubleLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseUniversalParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseUniversalParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterBigDecimalLiteral is called when production bigDecimalLiteral is entered.
func (s *BaseUniversalParserListener) EnterBigDecimalLiteral(ctx *BigDecimalLiteralContext) {}

// ExitBigDecimalLiteral is called when production bigDecimalLiteral is exited.
func (s *BaseUniversalParserListener) ExitBigDecimalLiteral(ctx *BigDecimalLiteralContext) {}

// EnterString is called when production string is entered.
func (s *BaseUniversalParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseUniversalParserListener) ExitString(ctx *StringContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseUniversalParserListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseUniversalParserListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseUniversalParserListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseUniversalParserListener) ExitNull(ctx *NullContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseUniversalParserListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseUniversalParserListener) ExitVersion(ctx *VersionContext) {}

// EnterNoReservedKeywords is called when production noReservedKeywords is entered.
func (s *BaseUniversalParserListener) EnterNoReservedKeywords(ctx *NoReservedKeywordsContext) {}

// ExitNoReservedKeywords is called when production noReservedKeywords is exited.
func (s *BaseUniversalParserListener) ExitNoReservedKeywords(ctx *NoReservedKeywordsContext) {}
