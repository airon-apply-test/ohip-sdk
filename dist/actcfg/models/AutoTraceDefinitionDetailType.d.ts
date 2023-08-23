/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AutoTraceBlockCategoryType } from './AutoTraceBlockCategoryType';
import type { AutoTraceGroupConfigType } from './AutoTraceGroupConfigType';
import type { AutoTraceId } from './AutoTraceId';
import type { AutoTraceTriggerCategoryType } from './AutoTraceTriggerCategoryType';
import type { ConditionGroupType } from './ConditionGroupType';
/**
 * Auto Trace Definition detail information.
 * @export
 * @interface AutoTraceDefinitionDetailType
 */
export interface AutoTraceDefinitionDetailType {
    /**
     *
     * @type {AutoTraceId}
     * @memberof AutoTraceDefinitionDetailType
     */
    autoTraceId?: AutoTraceId;
    /**
     * The hotel code where the Auto Trace should be created.
     * @type {string}
     * @memberof AutoTraceDefinitionDetailType
     */
    hotelId?: string;
    /**
     * Trace code
     * @type {string}
     * @memberof AutoTraceDefinitionDetailType
     */
    traceCode?: string;
    /**
     *
     * @type {AutoTraceGroupConfigType}
     * @memberof AutoTraceDefinitionDetailType
     */
    traceGroup?: AutoTraceGroupConfigType;
    /**
     *
     * @type {AutoTraceBlockCategoryType}
     * @memberof AutoTraceDefinitionDetailType
     */
    autoTraceBlockCategory?: AutoTraceBlockCategoryType;
    /**
     *
     * @type {AutoTraceTriggerCategoryType}
     * @memberof AutoTraceDefinitionDetailType
     */
    autoTraceTrigger?: AutoTraceTriggerCategoryType;
    /**
     * Element to hold column name of relavant table on selection of Update in AutoTraceTriggerCategoryType.
     * @type {string}
     * @memberof AutoTraceDefinitionDetailType
     */
    fieldNameOnUpdate?: string;
    /**
     * Element to hold column ID of relavant table on selection of Update in AutoTraceTriggerCategoryType.
     * @type {string}
     * @memberof AutoTraceDefinitionDetailType
     */
    fieldNameOnUpdateId?: string;
    /**
     *
     * @type {ConditionGroupType}
     * @memberof AutoTraceDefinitionDetailType
     */
    conditions?: ConditionGroupType;
}
/**
 * Check if a given object implements the AutoTraceDefinitionDetailType interface.
 */
export declare function instanceOfAutoTraceDefinitionDetailType(value: object): boolean;
export declare function AutoTraceDefinitionDetailTypeFromJSON(json: any): AutoTraceDefinitionDetailType;
export declare function AutoTraceDefinitionDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AutoTraceDefinitionDetailType;
export declare function AutoTraceDefinitionDetailTypeToJSON(value?: AutoTraceDefinitionDetailType | null): any;
