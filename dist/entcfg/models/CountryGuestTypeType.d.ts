/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { TranslationTextType2000 } from './TranslationTextType2000';
/**
 * Contains Common Master configuration detail.
 * @export
 * @interface CountryGuestTypeType
 */
export interface CountryGuestTypeType {
    /**
     * Common Master unique code.
     * @type {string}
     * @memberof CountryGuestTypeType
     */
    code?: string;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof CountryGuestTypeType
     */
    description?: TranslationTextType2000;
    /**
     * Common Master record sequence number.
     * @type {number}
     * @memberof CountryGuestTypeType
     */
    displayOrder?: number;
}
/**
 * Check if a given object implements the CountryGuestTypeType interface.
 */
export declare function instanceOfCountryGuestTypeType(value: object): boolean;
export declare function CountryGuestTypeTypeFromJSON(json: any): CountryGuestTypeType;
export declare function CountryGuestTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CountryGuestTypeType;
export declare function CountryGuestTypeTypeToJSON(value?: CountryGuestTypeType | null): any;
