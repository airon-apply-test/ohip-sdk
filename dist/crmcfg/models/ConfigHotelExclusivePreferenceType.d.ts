/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CodeDescriptionType } from './CodeDescriptionType';
/**
 * Base details common between both template and property level Exclusive preference ids.
 * @export
 * @interface ConfigHotelExclusivePreferenceType
 */
export interface ConfigHotelExclusivePreferenceType {
    /**
     * Specifies the Exclusive preference code.
     * @type {string}
     * @memberof ConfigHotelExclusivePreferenceType
     */
    code?: string;
    /**
     * Specifies the preference group the Exclusive preference belongs to.
     * @type {string}
     * @memberof ConfigHotelExclusivePreferenceType
     */
    preferenceGroup?: string;
    /**
     * Specifies the preference code and its description mapped to the exclusive preference.
     * @type {Array<CodeDescriptionType>}
     * @memberof ConfigHotelExclusivePreferenceType
     */
    preferenceCodes?: Array<CodeDescriptionType>;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ConfigHotelExclusivePreferenceType
     */
    orderSequence?: number;
    /**
     * Specifies the hotel code for which the Exclusive preference is specified.
     * @type {string}
     * @memberof ConfigHotelExclusivePreferenceType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the ConfigHotelExclusivePreferenceType interface.
 */
export declare function instanceOfConfigHotelExclusivePreferenceType(value: object): boolean;
export declare function ConfigHotelExclusivePreferenceTypeFromJSON(json: any): ConfigHotelExclusivePreferenceType;
export declare function ConfigHotelExclusivePreferenceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigHotelExclusivePreferenceType;
export declare function ConfigHotelExclusivePreferenceTypeToJSON(value?: ConfigHotelExclusivePreferenceType | null): any;
