/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
import type { AlienInfoType } from './AlienInfoType';
import {
    AlienInfoTypeFromJSON,
    AlienInfoTypeFromJSONTyped,
    AlienInfoTypeToJSON,
} from './AlienInfoType';
import type { AnonymizationType } from './AnonymizationType';
import {
    AnonymizationTypeFromJSON,
    AnonymizationTypeFromJSONTyped,
    AnonymizationTypeToJSON,
} from './AnonymizationType';
import type { CompanyInfoType } from './CompanyInfoType';
import {
    CompanyInfoTypeFromJSON,
    CompanyInfoTypeFromJSONTyped,
    CompanyInfoTypeToJSON,
} from './CompanyInfoType';
import type { CountryNameType } from './CountryNameType';
import {
    CountryNameTypeFromJSON,
    CountryNameTypeFromJSONTyped,
    CountryNameTypeToJSON,
} from './CountryNameType';
import type { CustomerTypeIdentifications } from './CustomerTypeIdentifications';
import {
    CustomerTypeIdentificationsFromJSON,
    CustomerTypeIdentificationsFromJSONTyped,
    CustomerTypeIdentificationsToJSON,
} from './CustomerTypeIdentifications';
import type { PersonNameType } from './PersonNameType';
import {
    PersonNameTypeFromJSON,
    PersonNameTypeFromJSONTyped,
    PersonNameTypeToJSON,
} from './PersonNameType';

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
export const CustomerTypeGenderEnum = {
    Male: 'Male',
    Female: 'Female',
    Unknown: 'Unknown'
} as const;
export type CustomerTypeGenderEnum = typeof CustomerTypeGenderEnum[keyof typeof CustomerTypeGenderEnum];


/**
 * Check if a given object implements the CustomerType interface.
 */
export function instanceOfCustomerType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CustomerTypeFromJSON(json: any): CustomerType {
    return CustomerTypeFromJSONTyped(json, false);
}

export function CustomerTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CustomerType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'personName': !exists(json, 'personName') ? undefined : ((json['personName'] as Array<any>).map(PersonNameTypeFromJSON)),
        'anonymization': !exists(json, 'anonymization') ? undefined : AnonymizationTypeFromJSON(json['anonymization']),
        'citizenCountry': !exists(json, 'citizenCountry') ? undefined : CountryNameTypeFromJSON(json['citizenCountry']),
        'identifications': !exists(json, 'identifications') ? undefined : CustomerTypeIdentificationsFromJSON(json['identifications']),
        'profession': !exists(json, 'profession') ? undefined : json['profession'],
        'alienInfo': !exists(json, 'alienInfo') ? undefined : AlienInfoTypeFromJSON(json['alienInfo']),
        'birthCountry': !exists(json, 'birthCountry') ? undefined : CountryNameTypeFromJSON(json['birthCountry']),
        'legalCompany': !exists(json, 'legalCompany') ? undefined : json['legalCompany'],
        'companyInfo': !exists(json, 'companyInfo') ? undefined : CompanyInfoTypeFromJSON(json['companyInfo']),
        'businessTitle': !exists(json, 'businessTitle') ? undefined : json['businessTitle'],
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'birthDate': !exists(json, 'birthDate') ? undefined : (new Date(json['birthDate'])),
        'birthDateMasked': !exists(json, 'birthDateMasked') ? undefined : json['birthDateMasked'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'language': !exists(json, 'language') ? undefined : json['language'],
        'nationality': !exists(json, 'nationality') ? undefined : json['nationality'],
        'nationalityDescription': !exists(json, 'nationalityDescription') ? undefined : json['nationalityDescription'],
        'customerValue': !exists(json, 'customerValue') ? undefined : json['customerValue'],
        'creditRating': !exists(json, 'creditRating') ? undefined : json['creditRating'],
        'vipStatus': !exists(json, 'vipStatus') ? undefined : json['vipStatus'],
        'vipDescription': !exists(json, 'vipDescription') ? undefined : json['vipDescription'],
        'birthPlace': !exists(json, 'birthPlace') ? undefined : json['birthPlace'],
        'privateProfile': !exists(json, 'privateProfile') ? undefined : json['privateProfile'],
        'blacklist': !exists(json, 'blacklist') ? undefined : json['blacklist'],
    };
}

export function CustomerTypeToJSON(value?: CustomerType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'personName': value.personName === undefined ? undefined : ((value.personName as Array<any>).map(PersonNameTypeToJSON)),
        'anonymization': AnonymizationTypeToJSON(value.anonymization),
        'citizenCountry': CountryNameTypeToJSON(value.citizenCountry),
        'identifications': CustomerTypeIdentificationsToJSON(value.identifications),
        'profession': value.profession,
        'alienInfo': AlienInfoTypeToJSON(value.alienInfo),
        'birthCountry': CountryNameTypeToJSON(value.birthCountry),
        'legalCompany': value.legalCompany,
        'companyInfo': CompanyInfoTypeToJSON(value.companyInfo),
        'businessTitle': value.businessTitle,
        'gender': value.gender,
        'birthDate': value.birthDate === undefined ? undefined : (value.birthDate.toISOString().substr(0,10)),
        'birthDateMasked': value.birthDateMasked,
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'language': value.language,
        'nationality': value.nationality,
        'nationalityDescription': value.nationalityDescription,
        'customerValue': value.customerValue,
        'creditRating': value.creditRating,
        'vipStatus': value.vipStatus,
        'vipDescription': value.vipDescription,
        'birthPlace': value.birthPlace,
        'privateProfile': value.privateProfile,
        'blacklist': value.blacklist,
    };
}

