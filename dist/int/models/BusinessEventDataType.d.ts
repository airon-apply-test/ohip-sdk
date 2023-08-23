/**
 * OPERA Cloud Integration Processor API
 * APIs to get Business Events generated in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BusinessEventIDType } from './BusinessEventIDType';
import type { BusinessEventType } from './BusinessEventType';
/**
 *
 * @export
 * @interface BusinessEventDataType
 */
export interface BusinessEventDataType {
    /**
     *
     * @type {BusinessEventType}
     * @memberof BusinessEventDataType
     */
    businessEvent?: BusinessEventType;
    /**
     *
     * @type {BusinessEventIDType}
     * @memberof BusinessEventDataType
     */
    businessEventId?: BusinessEventIDType;
}
/**
 * Check if a given object implements the BusinessEventDataType interface.
 */
export declare function instanceOfBusinessEventDataType(value: object): boolean;
export declare function BusinessEventDataTypeFromJSON(json: any): BusinessEventDataType;
export declare function BusinessEventDataTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BusinessEventDataType;
export declare function BusinessEventDataTypeToJSON(value?: BusinessEventDataType | null): any;
