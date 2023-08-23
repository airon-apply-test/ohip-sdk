/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ExpressionFieldType } from './ExpressionFieldType';
import type { ExpressionOperatorType } from './ExpressionOperatorType';
import type { ExpressionParameterType } from './ExpressionParameterType';
import type { LogicalOperatorType } from './LogicalOperatorType';
/**
 *
 * @export
 * @interface ConditionType
 */
export interface ConditionType {
    /**
     *
     * @type {ExpressionFieldType}
     * @memberof ConditionType
     */
    leftExpression?: ExpressionFieldType;
    /**
     *
     * @type {ExpressionOperatorType}
     * @memberof ConditionType
     */
    operator?: ExpressionOperatorType;
    /**
     *
     * @type {ExpressionParameterType}
     * @memberof ConditionType
     */
    rightExpression?: ExpressionParameterType;
    /**
     *
     * @type {LogicalOperatorType}
     * @memberof ConditionType
     */
    logicalOperator?: LogicalOperatorType;
}
/**
 * Check if a given object implements the ConditionType interface.
 */
export declare function instanceOfConditionType(value: object): boolean;
export declare function ConditionTypeFromJSON(json: any): ConditionType;
export declare function ConditionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConditionType;
export declare function ConditionTypeToJSON(value?: ConditionType | null): any;
