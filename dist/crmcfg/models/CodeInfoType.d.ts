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
import type { AddtionalCodeInfoTypeInner } from './AddtionalCodeInfoTypeInner';
/**
 *
 * @export
 * @interface CodeInfoType
 */
export interface CodeInfoType {
    /**
     *
     * @type {string}
     * @memberof CodeInfoType
     */
    description?: string;
    /**
     * Holds name of additional code information
     * @type {Array<AddtionalCodeInfoTypeInner>}
     * @memberof CodeInfoType
     */
    addtionalCodeInfo?: Array<AddtionalCodeInfoTypeInner>;
    /**
     *
     * @type {string}
     * @memberof CodeInfoType
     */
    hotelId?: string;
    /**
     *
     * @type {string}
     * @memberof CodeInfoType
     */
    code?: string;
}
/**
 * Check if a given object implements the CodeInfoType interface.
 */
export declare function instanceOfCodeInfoType(value: object): boolean;
export declare function CodeInfoTypeFromJSON(json: any): CodeInfoType;
export declare function CodeInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CodeInfoType;
export declare function CodeInfoTypeToJSON(value?: CodeInfoType | null): any;
