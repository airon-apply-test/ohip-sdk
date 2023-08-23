/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UniqueIDType } from './UniqueIDType';
/**
 *
 * @export
 * @interface MasterAccountInfoType
 */
export interface MasterAccountInfoType {
    /**
     *
     * @type {UniqueIDType}
     * @memberof MasterAccountInfoType
     */
    masterAccountId?: UniqueIDType;
    /**
     * Name of the Master account.
     * @type {string}
     * @memberof MasterAccountInfoType
     */
    masterAccountName?: string;
}
/**
 * Check if a given object implements the MasterAccountInfoType interface.
 */
export declare function instanceOfMasterAccountInfoType(value: object): boolean;
export declare function MasterAccountInfoTypeFromJSON(json: any): MasterAccountInfoType;
export declare function MasterAccountInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MasterAccountInfoType;
export declare function MasterAccountInfoTypeToJSON(value?: MasterAccountInfoType | null): any;
