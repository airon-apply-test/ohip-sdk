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
import type { CommonMasterColorType } from './CommonMasterColorType';
import type { TranslationTextType2000 } from './TranslationTextType2000';
/**
 * Contains Common Master configuration detail.
 * @export
 * @interface VIPLevelType
 */
export interface VIPLevelType {
    /**
     * Common Master unique code.
     * @type {string}
     * @memberof VIPLevelType
     */
    code?: string;
    /**
     *
     * @type {TranslationTextType2000}
     * @memberof VIPLevelType
     */
    description?: TranslationTextType2000;
    /**
     * Common Master record sequence number.
     * @type {number}
     * @memberof VIPLevelType
     */
    displayOrder?: number;
    /**
     *
     * @type {CommonMasterColorType}
     * @memberof VIPLevelType
     */
    displayColor?: CommonMasterColorType;
    /**
     * AI priority order of this code.
     * @type {number}
     * @memberof VIPLevelType
     */
    ranking?: number;
}
/**
 * Check if a given object implements the VIPLevelType interface.
 */
export declare function instanceOfVIPLevelType(value: object): boolean;
export declare function VIPLevelTypeFromJSON(json: any): VIPLevelType;
export declare function VIPLevelTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): VIPLevelType;
export declare function VIPLevelTypeToJSON(value?: VIPLevelType | null): any;
