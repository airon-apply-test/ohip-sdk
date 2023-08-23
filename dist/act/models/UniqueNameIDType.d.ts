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
/**
 * Defines descriptive name, unique identification, and basic information combination.
 * @export
 * @interface UniqueNameIDType
 */
export interface UniqueNameIDType {
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof UniqueNameIDType
     */
    primary?: boolean;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof UniqueNameIDType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof UniqueNameIDType
     */
    type?: string;
}
/**
 * Check if a given object implements the UniqueNameIDType interface.
 */
export declare function instanceOfUniqueNameIDType(value: object): boolean;
export declare function UniqueNameIDTypeFromJSON(json: any): UniqueNameIDType;
export declare function UniqueNameIDTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UniqueNameIDType;
export declare function UniqueNameIDTypeToJSON(value?: UniqueNameIDType | null): any;
