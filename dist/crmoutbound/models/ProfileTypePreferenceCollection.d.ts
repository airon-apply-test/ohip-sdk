/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PreferenceTypeType } from './PreferenceTypeType';
/**
 * List of customer preferences.
 * @export
 * @interface ProfileTypePreferenceCollection
 */
export interface ProfileTypePreferenceCollection {
    /**
     * Collection of Detailed information on preferences of the customer.
     * @type {Array<PreferenceTypeType>}
     * @memberof ProfileTypePreferenceCollection
     */
    preferenceType?: Array<PreferenceTypeType>;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypePreferenceCollection
     */
    totalRows?: number;
}
/**
 * Check if a given object implements the ProfileTypePreferenceCollection interface.
 */
export declare function instanceOfProfileTypePreferenceCollection(value: object): boolean;
export declare function ProfileTypePreferenceCollectionFromJSON(json: any): ProfileTypePreferenceCollection;
export declare function ProfileTypePreferenceCollectionFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypePreferenceCollection;
export declare function ProfileTypePreferenceCollectionToJSON(value?: ProfileTypePreferenceCollection | null): any;
