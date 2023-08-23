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
import type { CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions } from './CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions';
import type { UniqueIDType } from './UniqueIDType';
/**
 * The hotel code where the trace definition was created.
 * @export
 * @interface CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition
 */
export interface CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition {
    /**
     * The hotel code where the trace definition was created.
     * @type {string}
     * @memberof CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition
     */
    sourceHotelCode?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition
     */
    traceDefinitionCodeList?: Array<UniqueIDType>;
    /**
     *
     * @type {Array<string>}
     * @memberof CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition
     */
    targetHotelCode?: Array<string>;
    /**
     *
     * @type {CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions}
     * @memberof CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition
     */
    copyInstructions?: CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionCopyInstructions;
}
/**
 * Check if a given object implements the CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition interface.
 */
export declare function instanceOfCopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition(value: object): boolean;
export declare function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionFromJSON(json: any): CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition;
export declare function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionFromJSONTyped(json: any, ignoreDiscriminator: boolean): CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition;
export declare function CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinitionToJSON(value?: CopyAutoTraceDefinitionCriteriaTypeCopyAutoTraceDefinition | null): any;
