// Code generated from docs/UniversalParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // UniversalParser
import "github.com/antlr4-go/antlr/v4"

// UniversalParserListener is a complete listener for a parse tree produced by UniversalParser.
type UniversalParserListener interface {
	antlr.ParseTreeListener

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterViewTableAndView is called when entering the viewTableAndView production.
	EnterViewTableAndView(c *ViewTableAndViewContext)

	// EnterCreateMaterializedViewStatement is called when entering the createMaterializedViewStatement production.
	EnterCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// EnterUpdateMaterializedViewStatement is called when entering the updateMaterializedViewStatement production.
	EnterUpdateMaterializedViewStatement(c *UpdateMaterializedViewStatementContext)

	// EnterSetMaterializedViewLifecycleStatement is called when entering the setMaterializedViewLifecycleStatement production.
	EnterSetMaterializedViewLifecycleStatement(c *SetMaterializedViewLifecycleStatementContext)

	// EnterUpdateMaterializedViewLifecycleStatement is called when entering the updateMaterializedViewLifecycleStatement production.
	EnterUpdateMaterializedViewLifecycleStatement(c *UpdateMaterializedViewLifecycleStatementContext)

	// EnterDropMaterializedViewStatement is called when entering the dropMaterializedViewStatement production.
	EnterDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// EnterDropMaterializedViewPartitionStatement is called when entering the dropMaterializedViewPartitionStatement production.
	EnterDropMaterializedViewPartitionStatement(c *DropMaterializedViewPartitionStatementContext)

	// EnterCreateViewStatement is called when entering the createViewStatement production.
	EnterCreateViewStatement(c *CreateViewStatementContext)

	// EnterViewColumns is called when entering the viewColumns production.
	EnterViewColumns(c *ViewColumnsContext)

	// EnterRenameViewStatement is called when entering the renameViewStatement production.
	EnterRenameViewStatement(c *RenameViewStatementContext)

	// EnterChangeViewStatement is called when entering the changeViewStatement production.
	EnterChangeViewStatement(c *ChangeViewStatementContext)

	// EnterDropViewStatement is called when entering the dropViewStatement production.
	EnterDropViewStatement(c *DropViewStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterDropTableStatement is called when entering the dropTableStatement production.
	EnterDropTableStatement(c *DropTableStatementContext)

	// EnterCreateTableBodyspec is called when entering the createTableBodyspec production.
	EnterCreateTableBodyspec(c *CreateTableBodyspecContext)

	// EnterCreateTableItems is called when entering the createTableItems production.
	EnterCreateTableItems(c *CreateTableItemsContext)

	// EnterCreateTableItem is called when entering the createTableItem production.
	EnterCreateTableItem(c *CreateTableItemContext)

	// EnterTablePartitionSpec is called when entering the tablePartitionSpec production.
	EnterTablePartitionSpec(c *TablePartitionSpecContext)

	// EnterPartitionItems is called when entering the partitionItems production.
	EnterPartitionItems(c *PartitionItemsContext)

	// EnterPartitionItem is called when entering the partitionItem production.
	EnterPartitionItem(c *PartitionItemContext)

	// EnterSortSpec is called when entering the sortSpec production.
	EnterSortSpec(c *SortSpecContext)

	// EnterSkewedBySpec is called when entering the skewedBySpec production.
	EnterSkewedBySpec(c *SkewedBySpecContext)

	// EnterSkewedByColumns is called when entering the skewedByColumns production.
	EnterSkewedByColumns(c *SkewedByColumnsContext)

	// EnterSkewedByValues is called when entering the skewedByValues production.
	EnterSkewedByValues(c *SkewedByValuesContext)

	// EnterMultiSkewedValue is called when entering the multiSkewedValue production.
	EnterMultiSkewedValue(c *MultiSkewedValueContext)

	// EnterSimpleValue is called when entering the simpleValue production.
	EnterSimpleValue(c *SimpleValueContext)

	// EnterStorageDirectories is called when entering the storageDirectories production.
	EnterStorageDirectories(c *StorageDirectoriesContext)

	// EnterFormatSpec is called when entering the formatSpec production.
	EnterFormatSpec(c *FormatSpecContext)

	// EnterTableRowFormat is called when entering the tableRowFormat production.
	EnterTableRowFormat(c *TableRowFormatContext)

	// EnterStoredBy is called when entering the storedBy production.
	EnterStoredBy(c *StoredByContext)

	// EnterSerdeProperties is called when entering the serdeProperties production.
	EnterSerdeProperties(c *SerdePropertiesContext)

	// EnterStoredSpec is called when entering the storedSpec production.
	EnterStoredSpec(c *StoredSpecContext)

	// EnterFormatType is called when entering the formatType production.
	EnterFormatType(c *FormatTypeContext)

	// EnterPropertiesSpec is called when entering the propertiesSpec production.
	EnterPropertiesSpec(c *PropertiesSpecContext)

	// EnterTableSelectSpec is called when entering the tableSelectSpec production.
	EnterTableSelectSpec(c *TableSelectSpecContext)

	// EnterLifecycleSpec is called when entering the lifecycleSpec production.
	EnterLifecycleSpec(c *LifecycleSpecContext)

	// EnterInsertQueryStatement is called when entering the insertQueryStatement production.
	EnterInsertQueryStatement(c *InsertQueryStatementContext)

	// EnterInsertValueStatement is called when entering the insertValueStatement production.
	EnterInsertValueStatement(c *InsertValueStatementContext)

	// EnterDropTablePartitions is called when entering the dropTablePartitions production.
	EnterDropTablePartitions(c *DropTablePartitionsContext)

	// EnterAddTablePartitions is called when entering the addTablePartitions production.
	EnterAddTablePartitions(c *AddTablePartitionsContext)

	// EnterRenameTablePartitions is called when entering the renameTablePartitions production.
	EnterRenameTablePartitions(c *RenameTablePartitionsContext)

	// EnterInsertOverwriteTable is called when entering the insertOverwriteTable production.
	EnterInsertOverwriteTable(c *InsertOverwriteTableContext)

	// EnterInsertIntoTable is called when entering the insertIntoTable production.
	EnterInsertIntoTable(c *InsertIntoTableContext)

	// EnterInsertValues is called when entering the insertValues production.
	EnterInsertValues(c *InsertValuesContext)

	// EnterInsertTableColumns is called when entering the insertTableColumns production.
	EnterInsertTableColumns(c *InsertTableColumnsContext)

	// EnterPartitionSpec is called when entering the partitionSpec production.
	EnterPartitionSpec(c *PartitionSpecContext)

	// EnterStandardPartitionVal is called when entering the standardPartitionVal production.
	EnterStandardPartitionVal(c *StandardPartitionValContext)

	// EnterDefaultPartitionVal is called when entering the defaultPartitionVal production.
	EnterDefaultPartitionVal(c *DefaultPartitionValContext)

	// EnterIfOption is called when entering the ifOption production.
	EnterIfOption(c *IfOptionContext)

	// EnterReplaceOption is called when entering the replaceOption production.
	EnterReplaceOption(c *ReplaceOptionContext)

	// EnterNullOption is called when entering the nullOption production.
	EnterNullOption(c *NullOptionContext)

	// EnterDefaultOption is called when entering the defaultOption production.
	EnterDefaultOption(c *DefaultOptionContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterCte is called when entering the cte production.
	EnterCte(c *CteContext)

	// EnterNamedQuery is called when entering the namedQuery production.
	EnterNamedQuery(c *NamedQueryContext)

	// EnterColumnAliases is called when entering the columnAliases production.
	EnterColumnAliases(c *ColumnAliasesContext)

	// EnterQueryOrganization is called when entering the queryOrganization production.
	EnterQueryOrganization(c *QueryOrganizationContext)

	// EnterCombineOperation is called when entering the combineOperation production.
	EnterCombineOperation(c *CombineOperationContext)

	// EnterQueryTermDefault is called when entering the queryTermDefault production.
	EnterQueryTermDefault(c *QueryTermDefaultContext)

	// EnterQueryPrimaryDefault is called when entering the queryPrimaryDefault production.
	EnterQueryPrimaryDefault(c *QueryPrimaryDefaultContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterInlineTable is called when entering the inlineTable production.
	EnterInlineTable(c *InlineTableContext)

	// EnterFunctionTable is called when entering the functionTable production.
	EnterFunctionTable(c *FunctionTableContext)

	// EnterTemporalClause is called when entering the temporalClause production.
	EnterTemporalClause(c *TemporalClauseContext)

	// EnterSample is called when entering the sample production.
	EnterSample(c *SampleContext)

	// EnterSampleByPercentile is called when entering the sampleByPercentile production.
	EnterSampleByPercentile(c *SampleByPercentileContext)

	// EnterSampleByRows is called when entering the sampleByRows production.
	EnterSampleByRows(c *SampleByRowsContext)

	// EnterSampleByBucket is called when entering the sampleByBucket production.
	EnterSampleByBucket(c *SampleByBucketContext)

	// EnterSampleByBytes is called when entering the sampleByBytes production.
	EnterSampleByBytes(c *SampleByBytesContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterAliasedQuery is called when entering the aliasedQuery production.
	EnterAliasedQuery(c *AliasedQueryContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterRelationExtension is called when entering the relationExtension production.
	EnterRelationExtension(c *RelationExtensionContext)

	// EnterJoinRelation is called when entering the joinRelation production.
	EnterJoinRelation(c *JoinRelationContext)

	// EnterJoinCondtion is called when entering the joinCondtion production.
	EnterJoinCondtion(c *JoinCondtionContext)

	// EnterJoinType is called when entering the joinType production.
	EnterJoinType(c *JoinTypeContext)

	// EnterPivotClause is called when entering the pivotClause production.
	EnterPivotClause(c *PivotClauseContext)

	// EnterPivotAggregateItem is called when entering the pivotAggregateItem production.
	EnterPivotAggregateItem(c *PivotAggregateItemContext)

	// EnterPivotColumn is called when entering the pivotColumn production.
	EnterPivotColumn(c *PivotColumnContext)

	// EnterPivotValue is called when entering the pivotValue production.
	EnterPivotValue(c *PivotValueContext)

	// EnterUnpivotClause is called when entering the unpivotClause production.
	EnterUnpivotClause(c *UnpivotClauseContext)

	// EnterUnpivotOption is called when entering the unpivotOption production.
	EnterUnpivotOption(c *UnpivotOptionContext)

	// EnterUnpivotSingleValueColumn is called when entering the unpivotSingleValueColumn production.
	EnterUnpivotSingleValueColumn(c *UnpivotSingleValueColumnContext)

	// EnterUnpivotMultiValueColumn is called when entering the unpivotMultiValueColumn production.
	EnterUnpivotMultiValueColumn(c *UnpivotMultiValueColumnContext)

	// EnterSingleUnpivotNameColumn is called when entering the singleUnpivotNameColumn production.
	EnterSingleUnpivotNameColumn(c *SingleUnpivotNameColumnContext)

	// EnterMultiUnpivotNameColumn is called when entering the multiUnpivotNameColumn production.
	EnterMultiUnpivotNameColumn(c *MultiUnpivotNameColumnContext)

	// EnterUnpivotColumnItem is called when entering the unpivotColumnItem production.
	EnterUnpivotColumnItem(c *UnpivotColumnItemContext)

	// EnterUnpivotMultiColumnItem is called when entering the unpivotMultiColumnItem production.
	EnterUnpivotMultiColumnItem(c *UnpivotMultiColumnItemContext)

	// EnterUnpivotColumnAlias is called when entering the unpivotColumnAlias production.
	EnterUnpivotColumnAlias(c *UnpivotColumnAliasContext)

	// EnterUnpivotColumnAliasOption is called when entering the unpivotColumnAliasOption production.
	EnterUnpivotColumnAliasOption(c *UnpivotColumnAliasOptionContext)

	// EnterUnpivotColumnAliasValueList is called when entering the unpivotColumnAliasValueList production.
	EnterUnpivotColumnAliasValueList(c *UnpivotColumnAliasValueListContext)

	// EnterUnpivotColumnAliasValue is called when entering the unpivotColumnAliasValue production.
	EnterUnpivotColumnAliasValue(c *UnpivotColumnAliasValueContext)

	// EnterUnpivotColumn is called when entering the unpivotColumn production.
	EnterUnpivotColumn(c *UnpivotColumnContext)

	// EnterLateralView is called when entering the lateralView production.
	EnterLateralView(c *LateralViewContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterGroupExpr is called when entering the groupExpr production.
	EnterGroupExpr(c *GroupExprContext)

	// EnterGroupCollection is called when entering the groupCollection production.
	EnterGroupCollection(c *GroupCollectionContext)

	// EnterGroupWith is called when entering the groupWith production.
	EnterGroupWith(c *GroupWithContext)

	// EnterGroupWithSet is called when entering the groupWithSet production.
	EnterGroupWithSet(c *GroupWithSetContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupingElement is called when entering the groupingElement production.
	EnterGroupingElement(c *GroupingElementContext)

	// EnterGroupCube is called when entering the groupCube production.
	EnterGroupCube(c *GroupCubeContext)

	// EnterGroupSet is called when entering the groupSet production.
	EnterGroupSet(c *GroupSetContext)

	// EnterGroupingSet is called when entering the groupingSet production.
	EnterGroupingSet(c *GroupingSetContext)

	// EnterHint is called when entering the hint production.
	EnterHint(c *HintContext)

	// EnterHintStatement is called when entering the hintStatement production.
	EnterHintStatement(c *HintStatementContext)

	// EnterNamedExpressionSeq is called when entering the namedExpressionSeq production.
	EnterNamedExpressionSeq(c *NamedExpressionSeqContext)

	// EnterQueryStars is called when entering the queryStars production.
	EnterQueryStars(c *QueryStarsContext)

	// EnterQueryColumn is called when entering the queryColumn production.
	EnterQueryColumn(c *QueryColumnContext)

	// EnterAllQueryColumns is called when entering the allQueryColumns production.
	EnterAllQueryColumns(c *AllQueryColumnsContext)

	// EnterExpressionSeqs is called when entering the expressionSeqs production.
	EnterExpressionSeqs(c *ExpressionSeqsContext)

	// EnterExpressionSeq is called when entering the expressionSeq production.
	EnterExpressionSeq(c *ExpressionSeqContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalNot is called when entering the logicalNot production.
	EnterLogicalNot(c *LogicalNotContext)

	// EnterPredicated is called when entering the predicated production.
	EnterPredicated(c *PredicatedContext)

	// EnterExists is called when entering the exists production.
	EnterExists(c *ExistsContext)

	// EnterDefaultBinary is called when entering the defaultBinary production.
	EnterDefaultBinary(c *DefaultBinaryContext)

	// EnterBooleanPriority is called when entering the booleanPriority production.
	EnterBooleanPriority(c *BooleanPriorityContext)

	// EnterLogicalBinary is called when entering the logicalBinary production.
	EnterLogicalBinary(c *LogicalBinaryContext)

	// EnterBetweenPredicated is called when entering the betweenPredicated production.
	EnterBetweenPredicated(c *BetweenPredicatedContext)

	// EnterInValuePredicated is called when entering the inValuePredicated production.
	EnterInValuePredicated(c *InValuePredicatedContext)

	// EnterInQueryPredicated is called when entering the inQueryPredicated production.
	EnterInQueryPredicated(c *InQueryPredicatedContext)

	// EnterRlikePredicated is called when entering the rlikePredicated production.
	EnterRlikePredicated(c *RlikePredicatedContext)

	// EnterRegexpPredicated is called when entering the regexpPredicated production.
	EnterRegexpPredicated(c *RegexpPredicatedContext)

	// EnterLikeQuantifierPredicated is called when entering the likeQuantifierPredicated production.
	EnterLikeQuantifierPredicated(c *LikeQuantifierPredicatedContext)

	// EnterLikePatternPredicated is called when entering the likePatternPredicated production.
	EnterLikePatternPredicated(c *LikePatternPredicatedContext)

	// EnterNullPredicated is called when entering the nullPredicated production.
	EnterNullPredicated(c *NullPredicatedContext)

	// EnterBooleanPredicated is called when entering the booleanPredicated production.
	EnterBooleanPredicated(c *BooleanPredicatedContext)

	// EnterDistinctFromPredicated is called when entering the distinctFromPredicated production.
	EnterDistinctFromPredicated(c *DistinctFromPredicatedContext)

	// EnterDefaultUnary is called when entering the defaultUnary production.
	EnterDefaultUnary(c *DefaultUnaryContext)

	// EnterValuePriority is called when entering the valuePriority production.
	EnterValuePriority(c *ValuePriorityContext)

	// EnterArithmeticUnaryReverse is called when entering the arithmeticUnaryReverse production.
	EnterArithmeticUnaryReverse(c *ArithmeticUnaryReverseContext)

	// EnterArithmeticBinary is called when entering the arithmeticBinary production.
	EnterArithmeticBinary(c *ArithmeticBinaryContext)

	// EnterCompareBinary is called when entering the compareBinary production.
	EnterCompareBinary(c *CompareBinaryContext)

	// EnterDereference is called when entering the dereference production.
	EnterDereference(c *DereferenceContext)

	// EnterColumnReference is called when entering the columnReference production.
	EnterColumnReference(c *ColumnReferenceContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterFixValue is called when entering the fixValue production.
	EnterFixValue(c *FixValueContext)

	// EnterSubQueryExpression is called when entering the subQueryExpression production.
	EnterSubQueryExpression(c *SubQueryExpressionContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterSubScript is called when entering the subScript production.
	EnterSubScript(c *SubScriptContext)

	// EnterCaseExpress is called when entering the caseExpress production.
	EnterCaseExpress(c *CaseExpressContext)

	// EnterSearchedCase is called when entering the searchedCase production.
	EnterSearchedCase(c *SearchedCaseContext)

	// EnterSimpleCase is called when entering the simpleCase production.
	EnterSimpleCase(c *SimpleCaseContext)

	// EnterWhenClause is called when entering the whenClause production.
	EnterWhenClause(c *WhenClauseContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterSpecFunction is called when entering the specFunction production.
	EnterSpecFunction(c *SpecFunctionContext)

	// EnterCastFunction is called when entering the castFunction production.
	EnterCastFunction(c *CastFunctionContext)

	// EnterStructFunction is called when entering the structFunction production.
	EnterStructFunction(c *StructFunctionContext)

	// EnterFirstFunction is called when entering the firstFunction production.
	EnterFirstFunction(c *FirstFunctionContext)

	// EnterAnyValueFunction is called when entering the anyValueFunction production.
	EnterAnyValueFunction(c *AnyValueFunctionContext)

	// EnterLastFunction is called when entering the lastFunction production.
	EnterLastFunction(c *LastFunctionContext)

	// EnterSubQueryExpress is called when entering the subQueryExpress production.
	EnterSubQueryExpress(c *SubQueryExpressContext)

	// EnterTrimFunction is called when entering the trimFunction production.
	EnterTrimFunction(c *TrimFunctionContext)

	// EnterExtractFunction is called when entering the extractFunction production.
	EnterExtractFunction(c *ExtractFunctionContext)

	// EnterSubstringFunction is called when entering the substringFunction production.
	EnterSubstringFunction(c *SubstringFunctionContext)

	// EnterOverlayFunction is called when entering the overlayFunction production.
	EnterOverlayFunction(c *OverlayFunctionContext)

	// EnterPercentileFunction is called when entering the percentileFunction production.
	EnterPercentileFunction(c *PercentileFunctionContext)

	// EnterTransArrayFunction is called when entering the transArrayFunction production.
	EnterTransArrayFunction(c *TransArrayFunctionContext)

	// EnterStandardFunction is called when entering the standardFunction production.
	EnterStandardFunction(c *StandardFunctionContext)

	// EnterOverSpec is called when entering the overSpec production.
	EnterOverSpec(c *OverSpecContext)

	// EnterFilterSpec is called when entering the filterSpec production.
	EnterFilterSpec(c *FilterSpecContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterIntervalLiteral is called when entering the intervalLiteral production.
	EnterIntervalLiteral(c *IntervalLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterTypeLiteral is called when entering the typeLiteral production.
	EnterTypeLiteral(c *TypeLiteralContext)

	// EnterBooleaniteral is called when entering the booleaniteral production.
	EnterBooleaniteral(c *BooleaniteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterCompareOperator is called when entering the compareOperator production.
	EnterCompareOperator(c *CompareOperatorContext)

	// EnterArrayDataType is called when entering the arrayDataType production.
	EnterArrayDataType(c *ArrayDataTypeContext)

	// EnterMapDataType is called when entering the mapDataType production.
	EnterMapDataType(c *MapDataTypeContext)

	// EnterStructDataType is called when entering the structDataType production.
	EnterStructDataType(c *StructDataTypeContext)

	// EnterIntervalDataType is called when entering the intervalDataType production.
	EnterIntervalDataType(c *IntervalDataTypeContext)

	// EnterPrimitiveDataType is called when entering the primitiveDataType production.
	EnterPrimitiveDataType(c *PrimitiveDataTypeContext)

	// EnterComplexColTypeList is called when entering the complexColTypeList production.
	EnterComplexColTypeList(c *ComplexColTypeListContext)

	// EnterComplexColType is called when entering the complexColType production.
	EnterComplexColType(c *ComplexColTypeContext)

	// EnterSetQuantifier is called when entering the setQuantifier production.
	EnterSetQuantifier(c *SetQuantifierContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterMultiUnitsInterval is called when entering the multiUnitsInterval production.
	EnterMultiUnitsInterval(c *MultiUnitsIntervalContext)

	// EnterUnitToUnitInterval is called when entering the unitToUnitInterval production.
	EnterUnitToUnitInterval(c *UnitToUnitIntervalContext)

	// EnterMultiUnits is called when entering the multiUnits production.
	EnterMultiUnits(c *MultiUnitsContext)

	// EnterUnitToUnit is called when entering the unitToUnit production.
	EnterUnitToUnit(c *UnitToUnitContext)

	// EnterCommentSpec is called when entering the commentSpec production.
	EnterCommentSpec(c *CommentSpecContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterNamedWindow is called when entering the namedWindow production.
	EnterNamedWindow(c *NamedWindowContext)

	// EnterWindowRef is called when entering the windowRef production.
	EnterWindowRef(c *WindowRefContext)

	// EnterWindowDef is called when entering the windowDef production.
	EnterWindowDef(c *WindowDefContext)

	// EnterStandardWindowFrame is called when entering the standardWindowFrame production.
	EnterStandardWindowFrame(c *StandardWindowFrameContext)

	// EnterBetweenWindowFrame is called when entering the betweenWindowFrame production.
	EnterBetweenWindowFrame(c *BetweenWindowFrameContext)

	// EnterUnboundedPreceding is called when entering the unboundedPreceding production.
	EnterUnboundedPreceding(c *UnboundedPrecedingContext)

	// EnterCurrentRow is called when entering the currentRow production.
	EnterCurrentRow(c *CurrentRowContext)

	// EnterPreceding is called when entering the preceding production.
	EnterPreceding(c *PrecedingContext)

	// EnterTableProvider is called when entering the tableProvider production.
	EnterTableProvider(c *TableProviderContext)

	// EnterCreateFileFormat is called when entering the createFileFormat production.
	EnterCreateFileFormat(c *CreateFileFormatContext)

	// EnterTableFileFormat is called when entering the tableFileFormat production.
	EnterTableFileFormat(c *TableFileFormatContext)

	// EnterGenericFileFormat is called when entering the genericFileFormat production.
	EnterGenericFileFormat(c *GenericFileFormatContext)

	// EnterStorageHandler is called when entering the storageHandler production.
	EnterStorageHandler(c *StorageHandlerContext)

	// EnterRowFormatSerde is called when entering the rowFormatSerde production.
	EnterRowFormatSerde(c *RowFormatSerdeContext)

	// EnterRowFormatDelimited is called when entering the rowFormatDelimited production.
	EnterRowFormatDelimited(c *RowFormatDelimitedContext)

	// EnterPropertyList is called when entering the propertyList production.
	EnterPropertyList(c *PropertyListContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterPropertyKey is called when entering the propertyKey production.
	EnterPropertyKey(c *PropertyKeyContext)

	// EnterPropertyValue is called when entering the propertyValue production.
	EnterPropertyValue(c *PropertyValueContext)

	// EnterStandardArgs is called when entering the standardArgs production.
	EnterStandardArgs(c *StandardArgsContext)

	// EnterStars is called when entering the stars production.
	EnterStars(c *StarsContext)

	// EnterOrderBy is called when entering the orderBy production.
	EnterOrderBy(c *OrderByContext)

	// EnterSortBy is called when entering the sortBy production.
	EnterSortBy(c *SortByContext)

	// EnterDistributeBy is called when entering the distributeBy production.
	EnterDistributeBy(c *DistributeByContext)

	// EnterPartitionBy is called when entering the partitionBy production.
	EnterPartitionBy(c *PartitionByContext)

	// EnterClusterBy is called when entering the clusterBy production.
	EnterClusterBy(c *ClusterByContext)

	// EnterBucketsOption is called when entering the bucketsOption production.
	EnterBucketsOption(c *BucketsOptionContext)

	// EnterLimitItem is called when entering the limitItem production.
	EnterLimitItem(c *LimitItemContext)

	// EnterOffsetItem is called when entering the offsetItem production.
	EnterOffsetItem(c *OffsetItemContext)

	// EnterSortItems is called when entering the sortItems production.
	EnterSortItems(c *SortItemsContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterSortType is called when entering the sortType production.
	EnterSortType(c *SortTypeContext)

	// EnterAliasSpec is called when entering the aliasSpec production.
	EnterAliasSpec(c *AliasSpecContext)

	// EnterLocationSpec is called when entering the locationSpec production.
	EnterLocationSpec(c *LocationSpecContext)

	// EnterMultiIdentifierSeq is called when entering the multiIdentifierSeq production.
	EnterMultiIdentifierSeq(c *MultiIdentifierSeqContext)

	// EnterMultiIdentifier is called when entering the multiIdentifier production.
	EnterMultiIdentifier(c *MultiIdentifierContext)

	// EnterIdentifierSeq is called when entering the identifierSeq production.
	EnterIdentifierSeq(c *IdentifierSeqContext)

	// EnterUnquotedIdentifierAlternative is called when entering the unquotedIdentifierAlternative production.
	EnterUnquotedIdentifierAlternative(c *UnquotedIdentifierAlternativeContext)

	// EnterQuotedIdentifierAlternative is called when entering the quotedIdentifierAlternative production.
	EnterQuotedIdentifierAlternative(c *QuotedIdentifierAlternativeContext)

	// EnterQuotedIdentifier is called when entering the quotedIdentifier production.
	EnterQuotedIdentifier(c *QuotedIdentifierContext)

	// EnterExponentLiteral is called when entering the exponentLiteral production.
	EnterExponentLiteral(c *ExponentLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterLegacyDecimalLiteral is called when entering the legacyDecimalLiteral production.
	EnterLegacyDecimalLiteral(c *LegacyDecimalLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterBigIntLiteral is called when entering the bigIntLiteral production.
	EnterBigIntLiteral(c *BigIntLiteralContext)

	// EnterSmallIntLiteral is called when entering the smallIntLiteral production.
	EnterSmallIntLiteral(c *SmallIntLiteralContext)

	// EnterTinyIntLiteral is called when entering the tinyIntLiteral production.
	EnterTinyIntLiteral(c *TinyIntLiteralContext)

	// EnterDoubleLiteral is called when entering the doubleLiteral production.
	EnterDoubleLiteral(c *DoubleLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterBigDecimalLiteral is called when entering the bigDecimalLiteral production.
	EnterBigDecimalLiteral(c *BigDecimalLiteralContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterNoReservedKeywords is called when entering the noReservedKeywords production.
	EnterNoReservedKeywords(c *NoReservedKeywordsContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitViewTableAndView is called when exiting the viewTableAndView production.
	ExitViewTableAndView(c *ViewTableAndViewContext)

	// ExitCreateMaterializedViewStatement is called when exiting the createMaterializedViewStatement production.
	ExitCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// ExitUpdateMaterializedViewStatement is called when exiting the updateMaterializedViewStatement production.
	ExitUpdateMaterializedViewStatement(c *UpdateMaterializedViewStatementContext)

	// ExitSetMaterializedViewLifecycleStatement is called when exiting the setMaterializedViewLifecycleStatement production.
	ExitSetMaterializedViewLifecycleStatement(c *SetMaterializedViewLifecycleStatementContext)

	// ExitUpdateMaterializedViewLifecycleStatement is called when exiting the updateMaterializedViewLifecycleStatement production.
	ExitUpdateMaterializedViewLifecycleStatement(c *UpdateMaterializedViewLifecycleStatementContext)

	// ExitDropMaterializedViewStatement is called when exiting the dropMaterializedViewStatement production.
	ExitDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// ExitDropMaterializedViewPartitionStatement is called when exiting the dropMaterializedViewPartitionStatement production.
	ExitDropMaterializedViewPartitionStatement(c *DropMaterializedViewPartitionStatementContext)

	// ExitCreateViewStatement is called when exiting the createViewStatement production.
	ExitCreateViewStatement(c *CreateViewStatementContext)

	// ExitViewColumns is called when exiting the viewColumns production.
	ExitViewColumns(c *ViewColumnsContext)

	// ExitRenameViewStatement is called when exiting the renameViewStatement production.
	ExitRenameViewStatement(c *RenameViewStatementContext)

	// ExitChangeViewStatement is called when exiting the changeViewStatement production.
	ExitChangeViewStatement(c *ChangeViewStatementContext)

	// ExitDropViewStatement is called when exiting the dropViewStatement production.
	ExitDropViewStatement(c *DropViewStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitDropTableStatement is called when exiting the dropTableStatement production.
	ExitDropTableStatement(c *DropTableStatementContext)

	// ExitCreateTableBodyspec is called when exiting the createTableBodyspec production.
	ExitCreateTableBodyspec(c *CreateTableBodyspecContext)

	// ExitCreateTableItems is called when exiting the createTableItems production.
	ExitCreateTableItems(c *CreateTableItemsContext)

	// ExitCreateTableItem is called when exiting the createTableItem production.
	ExitCreateTableItem(c *CreateTableItemContext)

	// ExitTablePartitionSpec is called when exiting the tablePartitionSpec production.
	ExitTablePartitionSpec(c *TablePartitionSpecContext)

	// ExitPartitionItems is called when exiting the partitionItems production.
	ExitPartitionItems(c *PartitionItemsContext)

	// ExitPartitionItem is called when exiting the partitionItem production.
	ExitPartitionItem(c *PartitionItemContext)

	// ExitSortSpec is called when exiting the sortSpec production.
	ExitSortSpec(c *SortSpecContext)

	// ExitSkewedBySpec is called when exiting the skewedBySpec production.
	ExitSkewedBySpec(c *SkewedBySpecContext)

	// ExitSkewedByColumns is called when exiting the skewedByColumns production.
	ExitSkewedByColumns(c *SkewedByColumnsContext)

	// ExitSkewedByValues is called when exiting the skewedByValues production.
	ExitSkewedByValues(c *SkewedByValuesContext)

	// ExitMultiSkewedValue is called when exiting the multiSkewedValue production.
	ExitMultiSkewedValue(c *MultiSkewedValueContext)

	// ExitSimpleValue is called when exiting the simpleValue production.
	ExitSimpleValue(c *SimpleValueContext)

	// ExitStorageDirectories is called when exiting the storageDirectories production.
	ExitStorageDirectories(c *StorageDirectoriesContext)

	// ExitFormatSpec is called when exiting the formatSpec production.
	ExitFormatSpec(c *FormatSpecContext)

	// ExitTableRowFormat is called when exiting the tableRowFormat production.
	ExitTableRowFormat(c *TableRowFormatContext)

	// ExitStoredBy is called when exiting the storedBy production.
	ExitStoredBy(c *StoredByContext)

	// ExitSerdeProperties is called when exiting the serdeProperties production.
	ExitSerdeProperties(c *SerdePropertiesContext)

	// ExitStoredSpec is called when exiting the storedSpec production.
	ExitStoredSpec(c *StoredSpecContext)

	// ExitFormatType is called when exiting the formatType production.
	ExitFormatType(c *FormatTypeContext)

	// ExitPropertiesSpec is called when exiting the propertiesSpec production.
	ExitPropertiesSpec(c *PropertiesSpecContext)

	// ExitTableSelectSpec is called when exiting the tableSelectSpec production.
	ExitTableSelectSpec(c *TableSelectSpecContext)

	// ExitLifecycleSpec is called when exiting the lifecycleSpec production.
	ExitLifecycleSpec(c *LifecycleSpecContext)

	// ExitInsertQueryStatement is called when exiting the insertQueryStatement production.
	ExitInsertQueryStatement(c *InsertQueryStatementContext)

	// ExitInsertValueStatement is called when exiting the insertValueStatement production.
	ExitInsertValueStatement(c *InsertValueStatementContext)

	// ExitDropTablePartitions is called when exiting the dropTablePartitions production.
	ExitDropTablePartitions(c *DropTablePartitionsContext)

	// ExitAddTablePartitions is called when exiting the addTablePartitions production.
	ExitAddTablePartitions(c *AddTablePartitionsContext)

	// ExitRenameTablePartitions is called when exiting the renameTablePartitions production.
	ExitRenameTablePartitions(c *RenameTablePartitionsContext)

	// ExitInsertOverwriteTable is called when exiting the insertOverwriteTable production.
	ExitInsertOverwriteTable(c *InsertOverwriteTableContext)

	// ExitInsertIntoTable is called when exiting the insertIntoTable production.
	ExitInsertIntoTable(c *InsertIntoTableContext)

	// ExitInsertValues is called when exiting the insertValues production.
	ExitInsertValues(c *InsertValuesContext)

	// ExitInsertTableColumns is called when exiting the insertTableColumns production.
	ExitInsertTableColumns(c *InsertTableColumnsContext)

	// ExitPartitionSpec is called when exiting the partitionSpec production.
	ExitPartitionSpec(c *PartitionSpecContext)

	// ExitStandardPartitionVal is called when exiting the standardPartitionVal production.
	ExitStandardPartitionVal(c *StandardPartitionValContext)

	// ExitDefaultPartitionVal is called when exiting the defaultPartitionVal production.
	ExitDefaultPartitionVal(c *DefaultPartitionValContext)

	// ExitIfOption is called when exiting the ifOption production.
	ExitIfOption(c *IfOptionContext)

	// ExitReplaceOption is called when exiting the replaceOption production.
	ExitReplaceOption(c *ReplaceOptionContext)

	// ExitNullOption is called when exiting the nullOption production.
	ExitNullOption(c *NullOptionContext)

	// ExitDefaultOption is called when exiting the defaultOption production.
	ExitDefaultOption(c *DefaultOptionContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitCte is called when exiting the cte production.
	ExitCte(c *CteContext)

	// ExitNamedQuery is called when exiting the namedQuery production.
	ExitNamedQuery(c *NamedQueryContext)

	// ExitColumnAliases is called when exiting the columnAliases production.
	ExitColumnAliases(c *ColumnAliasesContext)

	// ExitQueryOrganization is called when exiting the queryOrganization production.
	ExitQueryOrganization(c *QueryOrganizationContext)

	// ExitCombineOperation is called when exiting the combineOperation production.
	ExitCombineOperation(c *CombineOperationContext)

	// ExitQueryTermDefault is called when exiting the queryTermDefault production.
	ExitQueryTermDefault(c *QueryTermDefaultContext)

	// ExitQueryPrimaryDefault is called when exiting the queryPrimaryDefault production.
	ExitQueryPrimaryDefault(c *QueryPrimaryDefaultContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitInlineTable is called when exiting the inlineTable production.
	ExitInlineTable(c *InlineTableContext)

	// ExitFunctionTable is called when exiting the functionTable production.
	ExitFunctionTable(c *FunctionTableContext)

	// ExitTemporalClause is called when exiting the temporalClause production.
	ExitTemporalClause(c *TemporalClauseContext)

	// ExitSample is called when exiting the sample production.
	ExitSample(c *SampleContext)

	// ExitSampleByPercentile is called when exiting the sampleByPercentile production.
	ExitSampleByPercentile(c *SampleByPercentileContext)

	// ExitSampleByRows is called when exiting the sampleByRows production.
	ExitSampleByRows(c *SampleByRowsContext)

	// ExitSampleByBucket is called when exiting the sampleByBucket production.
	ExitSampleByBucket(c *SampleByBucketContext)

	// ExitSampleByBytes is called when exiting the sampleByBytes production.
	ExitSampleByBytes(c *SampleByBytesContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitAliasedQuery is called when exiting the aliasedQuery production.
	ExitAliasedQuery(c *AliasedQueryContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitRelationExtension is called when exiting the relationExtension production.
	ExitRelationExtension(c *RelationExtensionContext)

	// ExitJoinRelation is called when exiting the joinRelation production.
	ExitJoinRelation(c *JoinRelationContext)

	// ExitJoinCondtion is called when exiting the joinCondtion production.
	ExitJoinCondtion(c *JoinCondtionContext)

	// ExitJoinType is called when exiting the joinType production.
	ExitJoinType(c *JoinTypeContext)

	// ExitPivotClause is called when exiting the pivotClause production.
	ExitPivotClause(c *PivotClauseContext)

	// ExitPivotAggregateItem is called when exiting the pivotAggregateItem production.
	ExitPivotAggregateItem(c *PivotAggregateItemContext)

	// ExitPivotColumn is called when exiting the pivotColumn production.
	ExitPivotColumn(c *PivotColumnContext)

	// ExitPivotValue is called when exiting the pivotValue production.
	ExitPivotValue(c *PivotValueContext)

	// ExitUnpivotClause is called when exiting the unpivotClause production.
	ExitUnpivotClause(c *UnpivotClauseContext)

	// ExitUnpivotOption is called when exiting the unpivotOption production.
	ExitUnpivotOption(c *UnpivotOptionContext)

	// ExitUnpivotSingleValueColumn is called when exiting the unpivotSingleValueColumn production.
	ExitUnpivotSingleValueColumn(c *UnpivotSingleValueColumnContext)

	// ExitUnpivotMultiValueColumn is called when exiting the unpivotMultiValueColumn production.
	ExitUnpivotMultiValueColumn(c *UnpivotMultiValueColumnContext)

	// ExitSingleUnpivotNameColumn is called when exiting the singleUnpivotNameColumn production.
	ExitSingleUnpivotNameColumn(c *SingleUnpivotNameColumnContext)

	// ExitMultiUnpivotNameColumn is called when exiting the multiUnpivotNameColumn production.
	ExitMultiUnpivotNameColumn(c *MultiUnpivotNameColumnContext)

	// ExitUnpivotColumnItem is called when exiting the unpivotColumnItem production.
	ExitUnpivotColumnItem(c *UnpivotColumnItemContext)

	// ExitUnpivotMultiColumnItem is called when exiting the unpivotMultiColumnItem production.
	ExitUnpivotMultiColumnItem(c *UnpivotMultiColumnItemContext)

	// ExitUnpivotColumnAlias is called when exiting the unpivotColumnAlias production.
	ExitUnpivotColumnAlias(c *UnpivotColumnAliasContext)

	// ExitUnpivotColumnAliasOption is called when exiting the unpivotColumnAliasOption production.
	ExitUnpivotColumnAliasOption(c *UnpivotColumnAliasOptionContext)

	// ExitUnpivotColumnAliasValueList is called when exiting the unpivotColumnAliasValueList production.
	ExitUnpivotColumnAliasValueList(c *UnpivotColumnAliasValueListContext)

	// ExitUnpivotColumnAliasValue is called when exiting the unpivotColumnAliasValue production.
	ExitUnpivotColumnAliasValue(c *UnpivotColumnAliasValueContext)

	// ExitUnpivotColumn is called when exiting the unpivotColumn production.
	ExitUnpivotColumn(c *UnpivotColumnContext)

	// ExitLateralView is called when exiting the lateralView production.
	ExitLateralView(c *LateralViewContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitGroupExpr is called when exiting the groupExpr production.
	ExitGroupExpr(c *GroupExprContext)

	// ExitGroupCollection is called when exiting the groupCollection production.
	ExitGroupCollection(c *GroupCollectionContext)

	// ExitGroupWith is called when exiting the groupWith production.
	ExitGroupWith(c *GroupWithContext)

	// ExitGroupWithSet is called when exiting the groupWithSet production.
	ExitGroupWithSet(c *GroupWithSetContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupingElement is called when exiting the groupingElement production.
	ExitGroupingElement(c *GroupingElementContext)

	// ExitGroupCube is called when exiting the groupCube production.
	ExitGroupCube(c *GroupCubeContext)

	// ExitGroupSet is called when exiting the groupSet production.
	ExitGroupSet(c *GroupSetContext)

	// ExitGroupingSet is called when exiting the groupingSet production.
	ExitGroupingSet(c *GroupingSetContext)

	// ExitHint is called when exiting the hint production.
	ExitHint(c *HintContext)

	// ExitHintStatement is called when exiting the hintStatement production.
	ExitHintStatement(c *HintStatementContext)

	// ExitNamedExpressionSeq is called when exiting the namedExpressionSeq production.
	ExitNamedExpressionSeq(c *NamedExpressionSeqContext)

	// ExitQueryStars is called when exiting the queryStars production.
	ExitQueryStars(c *QueryStarsContext)

	// ExitQueryColumn is called when exiting the queryColumn production.
	ExitQueryColumn(c *QueryColumnContext)

	// ExitAllQueryColumns is called when exiting the allQueryColumns production.
	ExitAllQueryColumns(c *AllQueryColumnsContext)

	// ExitExpressionSeqs is called when exiting the expressionSeqs production.
	ExitExpressionSeqs(c *ExpressionSeqsContext)

	// ExitExpressionSeq is called when exiting the expressionSeq production.
	ExitExpressionSeq(c *ExpressionSeqContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalNot is called when exiting the logicalNot production.
	ExitLogicalNot(c *LogicalNotContext)

	// ExitPredicated is called when exiting the predicated production.
	ExitPredicated(c *PredicatedContext)

	// ExitExists is called when exiting the exists production.
	ExitExists(c *ExistsContext)

	// ExitDefaultBinary is called when exiting the defaultBinary production.
	ExitDefaultBinary(c *DefaultBinaryContext)

	// ExitBooleanPriority is called when exiting the booleanPriority production.
	ExitBooleanPriority(c *BooleanPriorityContext)

	// ExitLogicalBinary is called when exiting the logicalBinary production.
	ExitLogicalBinary(c *LogicalBinaryContext)

	// ExitBetweenPredicated is called when exiting the betweenPredicated production.
	ExitBetweenPredicated(c *BetweenPredicatedContext)

	// ExitInValuePredicated is called when exiting the inValuePredicated production.
	ExitInValuePredicated(c *InValuePredicatedContext)

	// ExitInQueryPredicated is called when exiting the inQueryPredicated production.
	ExitInQueryPredicated(c *InQueryPredicatedContext)

	// ExitRlikePredicated is called when exiting the rlikePredicated production.
	ExitRlikePredicated(c *RlikePredicatedContext)

	// ExitRegexpPredicated is called when exiting the regexpPredicated production.
	ExitRegexpPredicated(c *RegexpPredicatedContext)

	// ExitLikeQuantifierPredicated is called when exiting the likeQuantifierPredicated production.
	ExitLikeQuantifierPredicated(c *LikeQuantifierPredicatedContext)

	// ExitLikePatternPredicated is called when exiting the likePatternPredicated production.
	ExitLikePatternPredicated(c *LikePatternPredicatedContext)

	// ExitNullPredicated is called when exiting the nullPredicated production.
	ExitNullPredicated(c *NullPredicatedContext)

	// ExitBooleanPredicated is called when exiting the booleanPredicated production.
	ExitBooleanPredicated(c *BooleanPredicatedContext)

	// ExitDistinctFromPredicated is called when exiting the distinctFromPredicated production.
	ExitDistinctFromPredicated(c *DistinctFromPredicatedContext)

	// ExitDefaultUnary is called when exiting the defaultUnary production.
	ExitDefaultUnary(c *DefaultUnaryContext)

	// ExitValuePriority is called when exiting the valuePriority production.
	ExitValuePriority(c *ValuePriorityContext)

	// ExitArithmeticUnaryReverse is called when exiting the arithmeticUnaryReverse production.
	ExitArithmeticUnaryReverse(c *ArithmeticUnaryReverseContext)

	// ExitArithmeticBinary is called when exiting the arithmeticBinary production.
	ExitArithmeticBinary(c *ArithmeticBinaryContext)

	// ExitCompareBinary is called when exiting the compareBinary production.
	ExitCompareBinary(c *CompareBinaryContext)

	// ExitDereference is called when exiting the dereference production.
	ExitDereference(c *DereferenceContext)

	// ExitColumnReference is called when exiting the columnReference production.
	ExitColumnReference(c *ColumnReferenceContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitFixValue is called when exiting the fixValue production.
	ExitFixValue(c *FixValueContext)

	// ExitSubQueryExpression is called when exiting the subQueryExpression production.
	ExitSubQueryExpression(c *SubQueryExpressionContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitSubScript is called when exiting the subScript production.
	ExitSubScript(c *SubScriptContext)

	// ExitCaseExpress is called when exiting the caseExpress production.
	ExitCaseExpress(c *CaseExpressContext)

	// ExitSearchedCase is called when exiting the searchedCase production.
	ExitSearchedCase(c *SearchedCaseContext)

	// ExitSimpleCase is called when exiting the simpleCase production.
	ExitSimpleCase(c *SimpleCaseContext)

	// ExitWhenClause is called when exiting the whenClause production.
	ExitWhenClause(c *WhenClauseContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitSpecFunction is called when exiting the specFunction production.
	ExitSpecFunction(c *SpecFunctionContext)

	// ExitCastFunction is called when exiting the castFunction production.
	ExitCastFunction(c *CastFunctionContext)

	// ExitStructFunction is called when exiting the structFunction production.
	ExitStructFunction(c *StructFunctionContext)

	// ExitFirstFunction is called when exiting the firstFunction production.
	ExitFirstFunction(c *FirstFunctionContext)

	// ExitAnyValueFunction is called when exiting the anyValueFunction production.
	ExitAnyValueFunction(c *AnyValueFunctionContext)

	// ExitLastFunction is called when exiting the lastFunction production.
	ExitLastFunction(c *LastFunctionContext)

	// ExitSubQueryExpress is called when exiting the subQueryExpress production.
	ExitSubQueryExpress(c *SubQueryExpressContext)

	// ExitTrimFunction is called when exiting the trimFunction production.
	ExitTrimFunction(c *TrimFunctionContext)

	// ExitExtractFunction is called when exiting the extractFunction production.
	ExitExtractFunction(c *ExtractFunctionContext)

	// ExitSubstringFunction is called when exiting the substringFunction production.
	ExitSubstringFunction(c *SubstringFunctionContext)

	// ExitOverlayFunction is called when exiting the overlayFunction production.
	ExitOverlayFunction(c *OverlayFunctionContext)

	// ExitPercentileFunction is called when exiting the percentileFunction production.
	ExitPercentileFunction(c *PercentileFunctionContext)

	// ExitTransArrayFunction is called when exiting the transArrayFunction production.
	ExitTransArrayFunction(c *TransArrayFunctionContext)

	// ExitStandardFunction is called when exiting the standardFunction production.
	ExitStandardFunction(c *StandardFunctionContext)

	// ExitOverSpec is called when exiting the overSpec production.
	ExitOverSpec(c *OverSpecContext)

	// ExitFilterSpec is called when exiting the filterSpec production.
	ExitFilterSpec(c *FilterSpecContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitIntervalLiteral is called when exiting the intervalLiteral production.
	ExitIntervalLiteral(c *IntervalLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitTypeLiteral is called when exiting the typeLiteral production.
	ExitTypeLiteral(c *TypeLiteralContext)

	// ExitBooleaniteral is called when exiting the booleaniteral production.
	ExitBooleaniteral(c *BooleaniteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitCompareOperator is called when exiting the compareOperator production.
	ExitCompareOperator(c *CompareOperatorContext)

	// ExitArrayDataType is called when exiting the arrayDataType production.
	ExitArrayDataType(c *ArrayDataTypeContext)

	// ExitMapDataType is called when exiting the mapDataType production.
	ExitMapDataType(c *MapDataTypeContext)

	// ExitStructDataType is called when exiting the structDataType production.
	ExitStructDataType(c *StructDataTypeContext)

	// ExitIntervalDataType is called when exiting the intervalDataType production.
	ExitIntervalDataType(c *IntervalDataTypeContext)

	// ExitPrimitiveDataType is called when exiting the primitiveDataType production.
	ExitPrimitiveDataType(c *PrimitiveDataTypeContext)

	// ExitComplexColTypeList is called when exiting the complexColTypeList production.
	ExitComplexColTypeList(c *ComplexColTypeListContext)

	// ExitComplexColType is called when exiting the complexColType production.
	ExitComplexColType(c *ComplexColTypeContext)

	// ExitSetQuantifier is called when exiting the setQuantifier production.
	ExitSetQuantifier(c *SetQuantifierContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitMultiUnitsInterval is called when exiting the multiUnitsInterval production.
	ExitMultiUnitsInterval(c *MultiUnitsIntervalContext)

	// ExitUnitToUnitInterval is called when exiting the unitToUnitInterval production.
	ExitUnitToUnitInterval(c *UnitToUnitIntervalContext)

	// ExitMultiUnits is called when exiting the multiUnits production.
	ExitMultiUnits(c *MultiUnitsContext)

	// ExitUnitToUnit is called when exiting the unitToUnit production.
	ExitUnitToUnit(c *UnitToUnitContext)

	// ExitCommentSpec is called when exiting the commentSpec production.
	ExitCommentSpec(c *CommentSpecContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitNamedWindow is called when exiting the namedWindow production.
	ExitNamedWindow(c *NamedWindowContext)

	// ExitWindowRef is called when exiting the windowRef production.
	ExitWindowRef(c *WindowRefContext)

	// ExitWindowDef is called when exiting the windowDef production.
	ExitWindowDef(c *WindowDefContext)

	// ExitStandardWindowFrame is called when exiting the standardWindowFrame production.
	ExitStandardWindowFrame(c *StandardWindowFrameContext)

	// ExitBetweenWindowFrame is called when exiting the betweenWindowFrame production.
	ExitBetweenWindowFrame(c *BetweenWindowFrameContext)

	// ExitUnboundedPreceding is called when exiting the unboundedPreceding production.
	ExitUnboundedPreceding(c *UnboundedPrecedingContext)

	// ExitCurrentRow is called when exiting the currentRow production.
	ExitCurrentRow(c *CurrentRowContext)

	// ExitPreceding is called when exiting the preceding production.
	ExitPreceding(c *PrecedingContext)

	// ExitTableProvider is called when exiting the tableProvider production.
	ExitTableProvider(c *TableProviderContext)

	// ExitCreateFileFormat is called when exiting the createFileFormat production.
	ExitCreateFileFormat(c *CreateFileFormatContext)

	// ExitTableFileFormat is called when exiting the tableFileFormat production.
	ExitTableFileFormat(c *TableFileFormatContext)

	// ExitGenericFileFormat is called when exiting the genericFileFormat production.
	ExitGenericFileFormat(c *GenericFileFormatContext)

	// ExitStorageHandler is called when exiting the storageHandler production.
	ExitStorageHandler(c *StorageHandlerContext)

	// ExitRowFormatSerde is called when exiting the rowFormatSerde production.
	ExitRowFormatSerde(c *RowFormatSerdeContext)

	// ExitRowFormatDelimited is called when exiting the rowFormatDelimited production.
	ExitRowFormatDelimited(c *RowFormatDelimitedContext)

	// ExitPropertyList is called when exiting the propertyList production.
	ExitPropertyList(c *PropertyListContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitPropertyKey is called when exiting the propertyKey production.
	ExitPropertyKey(c *PropertyKeyContext)

	// ExitPropertyValue is called when exiting the propertyValue production.
	ExitPropertyValue(c *PropertyValueContext)

	// ExitStandardArgs is called when exiting the standardArgs production.
	ExitStandardArgs(c *StandardArgsContext)

	// ExitStars is called when exiting the stars production.
	ExitStars(c *StarsContext)

	// ExitOrderBy is called when exiting the orderBy production.
	ExitOrderBy(c *OrderByContext)

	// ExitSortBy is called when exiting the sortBy production.
	ExitSortBy(c *SortByContext)

	// ExitDistributeBy is called when exiting the distributeBy production.
	ExitDistributeBy(c *DistributeByContext)

	// ExitPartitionBy is called when exiting the partitionBy production.
	ExitPartitionBy(c *PartitionByContext)

	// ExitClusterBy is called when exiting the clusterBy production.
	ExitClusterBy(c *ClusterByContext)

	// ExitBucketsOption is called when exiting the bucketsOption production.
	ExitBucketsOption(c *BucketsOptionContext)

	// ExitLimitItem is called when exiting the limitItem production.
	ExitLimitItem(c *LimitItemContext)

	// ExitOffsetItem is called when exiting the offsetItem production.
	ExitOffsetItem(c *OffsetItemContext)

	// ExitSortItems is called when exiting the sortItems production.
	ExitSortItems(c *SortItemsContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitSortType is called when exiting the sortType production.
	ExitSortType(c *SortTypeContext)

	// ExitAliasSpec is called when exiting the aliasSpec production.
	ExitAliasSpec(c *AliasSpecContext)

	// ExitLocationSpec is called when exiting the locationSpec production.
	ExitLocationSpec(c *LocationSpecContext)

	// ExitMultiIdentifierSeq is called when exiting the multiIdentifierSeq production.
	ExitMultiIdentifierSeq(c *MultiIdentifierSeqContext)

	// ExitMultiIdentifier is called when exiting the multiIdentifier production.
	ExitMultiIdentifier(c *MultiIdentifierContext)

	// ExitIdentifierSeq is called when exiting the identifierSeq production.
	ExitIdentifierSeq(c *IdentifierSeqContext)

	// ExitUnquotedIdentifierAlternative is called when exiting the unquotedIdentifierAlternative production.
	ExitUnquotedIdentifierAlternative(c *UnquotedIdentifierAlternativeContext)

	// ExitQuotedIdentifierAlternative is called when exiting the quotedIdentifierAlternative production.
	ExitQuotedIdentifierAlternative(c *QuotedIdentifierAlternativeContext)

	// ExitQuotedIdentifier is called when exiting the quotedIdentifier production.
	ExitQuotedIdentifier(c *QuotedIdentifierContext)

	// ExitExponentLiteral is called when exiting the exponentLiteral production.
	ExitExponentLiteral(c *ExponentLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitLegacyDecimalLiteral is called when exiting the legacyDecimalLiteral production.
	ExitLegacyDecimalLiteral(c *LegacyDecimalLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitBigIntLiteral is called when exiting the bigIntLiteral production.
	ExitBigIntLiteral(c *BigIntLiteralContext)

	// ExitSmallIntLiteral is called when exiting the smallIntLiteral production.
	ExitSmallIntLiteral(c *SmallIntLiteralContext)

	// ExitTinyIntLiteral is called when exiting the tinyIntLiteral production.
	ExitTinyIntLiteral(c *TinyIntLiteralContext)

	// ExitDoubleLiteral is called when exiting the doubleLiteral production.
	ExitDoubleLiteral(c *DoubleLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitBigDecimalLiteral is called when exiting the bigDecimalLiteral production.
	ExitBigDecimalLiteral(c *BigDecimalLiteralContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitNoReservedKeywords is called when exiting the noReservedKeywords production.
	ExitNoReservedKeywords(c *NoReservedKeywordsContext)
}
