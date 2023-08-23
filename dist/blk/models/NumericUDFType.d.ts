/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Used to hold user defined field of Numeric Type. It is highly recommended to use UDFN01, UDFN02,...UDFN40 (Total 40) as Numeric UDF names(commonly used on Reservation, Profile etc.). Name is not restricted using enumeration, to provide flexibility of different name usage if required.
 * @export
 * @interface NumericUDFType
 */
export interface NumericUDFType {
    /**
     * Name of user defined field.
     * @type {string}
     * @memberof NumericUDFType
     */
    name?: string;
    /**
     * Value of user defined field.
     * @type {number}
     * @memberof NumericUDFType
     */
    value?: number;
    /**
     * Label of user defined field used by vendors or customers.
     * @type {string}
     * @memberof NumericUDFType
     */
    alternateName?: string;
}
/**
 * Check if a given object implements the NumericUDFType interface.
 */
export declare function instanceOfNumericUDFType(value: object): boolean;
export declare function NumericUDFTypeFromJSON(json: any): NumericUDFType;
export declare function NumericUDFTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NumericUDFType;
export declare function NumericUDFTypeToJSON(value?: NumericUDFType | null): any;
