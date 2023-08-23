/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { NameValueType } from './NameValueType';
/**
 * Detailed information returned during name value search.
 * @export
 * @interface NameValueDetailType
 */
export interface NameValueDetailType {
    /**
     * List of name value pairs.
     * @type {Array<NameValueType>}
     * @memberof NameValueDetailType
     */
    nameValues?: Array<NameValueType>;
}
/**
 * Check if a given object implements the NameValueDetailType interface.
 */
export declare function instanceOfNameValueDetailType(value: object): boolean;
export declare function NameValueDetailTypeFromJSON(json: any): NameValueDetailType;
export declare function NameValueDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NameValueDetailType;
export declare function NameValueDetailTypeToJSON(value?: NameValueDetailType | null): any;
