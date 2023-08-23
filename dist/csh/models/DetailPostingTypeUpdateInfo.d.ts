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
/**
 * Update info associated to this transaction.
 * @export
 * @interface DetailPostingTypeUpdateInfo
 */
export interface DetailPostingTypeUpdateInfo {
    /**
     *
     * @type {Date}
     * @memberof DetailPostingTypeUpdateInfo
     */
    updateDate?: Date;
    /**
     *
     * @type {string}
     * @memberof DetailPostingTypeUpdateInfo
     */
    updateBy?: string;
}
/**
 * Check if a given object implements the DetailPostingTypeUpdateInfo interface.
 */
export declare function instanceOfDetailPostingTypeUpdateInfo(value: object): boolean;
export declare function DetailPostingTypeUpdateInfoFromJSON(json: any): DetailPostingTypeUpdateInfo;
export declare function DetailPostingTypeUpdateInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): DetailPostingTypeUpdateInfo;
export declare function DetailPostingTypeUpdateInfoToJSON(value?: DetailPostingTypeUpdateInfo | null): any;
