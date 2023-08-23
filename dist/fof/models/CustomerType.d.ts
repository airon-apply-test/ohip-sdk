/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AlienInfoType } from './AlienInfoType';
import type { AnonymizationType } from './AnonymizationType';
import type { CompanyInfoType } from './CompanyInfoType';
import type { CountryNameType } from './CountryNameType';
import type { CustomerTypeIdentifications } from './CustomerTypeIdentifications';
import type { PersonNameType } from './PersonNameType';
/**
 * Contains basic data on the customer's identity, location, relationships, finances, memberships, etc.
 * @export
 * @interface CustomerType
 */
export interface CustomerType {
    /**
     * Detailed name information for the customer.
     * @type {Array<PersonNameType>}
     * @memberof CustomerType
     */
    personName?: Array<PersonNameType>;
    /**
     *
     * @type {AnonymizationType}
     * @memberof CustomerType
     */
    anonymization?: AnonymizationType;
    /**
     *
     * @type {CountryNameType}
     * @memberof CustomerType
     */
    citizenCountry?: CountryNameType;
    /**
     *
     * @type {CustomerTypeIdentifications}
     * @memberof CustomerType
     */
    identifications?: CustomerTypeIdentifications;
    /**
     * Profession of a person.
     * @type {string}
     * @memberof CustomerType
     */
    profession?: string;
    /**
     *
     * @type {AlienInfoType}
     * @memberof CustomerType
     */
    alienInfo?: AlienInfoType;
    /**
     *
     * @type {CountryNameType}
     * @memberof CustomerType
     */
    birthCountry?: CountryNameType;
    /**
     * Name Of the company the individual is associated with.
     * @type {string}
     * @memberof CustomerType
     */
    legalCompany?: string;
    /**
     *
     * @type {CompanyInfoType}
     * @memberof CustomerType
     */
    companyInfo?: CompanyInfoType;
    /**
     * Business Title.
     * @type {string}
     * @memberof CustomerType
     */
    businessTitle?: string;
    /**
     * Identifies the gender.
     * @type {string}
     * @memberof CustomerType
     */
    gender?: CustomerTypeGenderEnum;
    /**
     * Indicates the date of birth as indicated in the document, in ISO 8601 prescribed format.
     * @type {Date}
     * @memberof CustomerType
     */
    birthDate?: Date;
    /**
     * Indicates the date of birth as masked.
     * @type {string}
     * @memberof CustomerType
     */
    birthDateMasked?: string;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof CustomerType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof CustomerType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof CustomerType
     */
    decimalPlaces?: number;
    /**
     * Language identification.
     * @type {string}
     * @memberof CustomerType
     */
    language?: string;
    /**
     * Nationality code identification
     * @type {string}
     * @memberof CustomerType
     */
    nationality?: string;
    /**
     * Nationality code description
     * @type {string}
     * @memberof CustomerType
     */
    nationalityDescription?: string;
    /**
     * The supplier's ranking of the customer (e.g., VIP, numerical ranking).
     * @type {string}
     * @memberof CustomerType
     */
    customerValue?: string;
    /**
     * Credit Rating of the customer.
     * @type {string}
     * @memberof CustomerType
     */
    creditRating?: string;
    /**
     * VIP status of the customer.
     * @type {string}
     * @memberof CustomerType
     */
    vipStatus?: string;
    /**
     * Description of the VIP status.
     * @type {string}
     * @memberof CustomerType
     */
    vipDescription?: string;
    /**
     * Place of birth.
     * @type {string}
     * @memberof CustomerType
     */
    birthPlace?: string;
    /**
     * This element tells profile is property exclusive or not.
     * @type {boolean}
     * @memberof CustomerType
     */
    privateProfile?: boolean;
    /**
     * This element tells if profile is blacklisted or not.
     * @type {boolean}
     * @memberof CustomerType
     */
    blacklist?: boolean;
}
/**
 * @export
 */
export declare const CustomerTypeGenderEnum: {
    readonly Male: "Male";
    readonly Female: "Female";
    readonly Unknown: "Unknown";
};
export type CustomerTypeGenderEnum = typeof CustomerTypeGenderEnum[keyof typeof CustomerTypeGenderEnum];
/**
 * Check if a given object implements the CustomerType interface.
 */
export declare function instanceOfCustomerType(value: object): boolean;
export declare function CustomerTypeFromJSON(json: any): CustomerType;
export declare function CustomerTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CustomerType;
export declare function CustomerTypeToJSON(value?: CustomerType | null): any;
