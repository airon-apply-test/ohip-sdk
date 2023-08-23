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
 * Identification information of the customer.
 * @export
 * @interface IdentificationType
 */
export interface IdentificationType {
    /**
     * Identification Type. Eg Passport, Driving License etc.
     * @type {string}
     * @memberof IdentificationType
     */
    idType?: string;
    /**
     * Identification Number.
     * @type {string}
     * @memberof IdentificationType
     */
    idNumber?: string;
    /**
     * Masked Identification Number.
     * @type {string}
     * @memberof IdentificationType
     */
    idNumberMasked?: string;
    /**
     * The country where Identification was issued.
     * @type {string}
     * @memberof IdentificationType
     */
    issuedCountry?: string;
    /**
     * Issued date of Identification.
     * @type {Date}
     * @memberof IdentificationType
     */
    issuedDate?: Date;
    /**
     * The place where Identification was issued.
     * @type {string}
     * @memberof IdentificationType
     */
    issuedPlace?: string;
    /**
     * Expiration date of Identification.
     * @type {Date}
     * @memberof IdentificationType
     */
    expirationDate?: Date;
    /**
     * Property where the identification belongs to.
     * @type {string}
     * @memberof IdentificationType
     */
    registeredProperty?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof IdentificationType
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof IdentificationType
     */
    orderSequence?: number;
}
/**
 * Check if a given object implements the IdentificationType interface.
 */
export declare function instanceOfIdentificationType(value: object): boolean;
export declare function IdentificationTypeFromJSON(json: any): IdentificationType;
export declare function IdentificationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdentificationType;
export declare function IdentificationTypeToJSON(value?: IdentificationType | null): any;
