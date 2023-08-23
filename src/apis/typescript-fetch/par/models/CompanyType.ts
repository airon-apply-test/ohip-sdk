/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { IATAInfoType } from './IATAInfoType';
import {
    IATAInfoTypeFromJSON,
    IATAInfoTypeFromJSONTyped,
    IATAInfoTypeToJSON,
} from './IATAInfoType';

/**
 * 
 * @export
 * @interface CompanyType
 */
export interface CompanyType {
    /**
     * Name of the company.
     * @type {string}
     * @memberof CompanyType
     */
    companyName?: string;
    /**
     * Second Name of the company.
     * @type {string}
     * @memberof CompanyType
     */
    companyName2?: string;
    /**
     * Third Name of the company.
     * @type {string}
     * @memberof CompanyType
     */
    companyName3?: string;
    /**
     * Alternate Name of the Company. Mainly, it's the name of the company written in the Alternate Language.
     * @type {string}
     * @memberof CompanyType
     */
    alternateName?: string;
    /**
     * Alternate Language of the company.
     * @type {string}
     * @memberof CompanyType
     */
    alternateLanguage?: string;
    /**
     * A jurisdiction(Territory) in which a company is authorized to do business.
     * @type {string}
     * @memberof CompanyType
     */
    businessLocale?: string;
    /**
     * 
     * @type {IATAInfoType}
     * @memberof CompanyType
     */
    iATAInfo?: IATAInfoType;
    /**
     * Business Title.
     * @type {string}
     * @memberof CompanyType
     */
    businessTitle?: string;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof CompanyType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof CompanyType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof CompanyType
     */
    decimalPlaces?: number;
    /**
     * Language identification.
     * @type {string}
     * @memberof CompanyType
     */
    language?: string;
    /**
     * Nationality code identification
     * @type {string}
     * @memberof CompanyType
     */
    nationality?: string;
    /**
     * Nationality code description
     * @type {string}
     * @memberof CompanyType
     */
    nationalityDescription?: string;
    /**
     * Commission code of the company.
     * @type {string}
     * @memberof CompanyType
     */
    commissionCode?: string;
    /**
     * Credit Rating of the company.
     * @type {string}
     * @memberof CompanyType
     */
    creditRating?: string;
    /**
     * The type of corporate ID
     * @type {string}
     * @memberof CompanyType
     */
    corporateIdType?: string;
    /**
     * VIP status of the company.
     * @type {string}
     * @memberof CompanyType
     */
    vipStatus?: string;
    /**
     * Description of the VIP status.
     * @type {string}
     * @memberof CompanyType
     */
    vipDescription?: string;
    /**
     * This element tells if profile is blacklisted or not.
     * @type {boolean}
     * @memberof CompanyType
     */
    blacklist?: boolean;
}

/**
 * Check if a given object implements the CompanyType interface.
 */
export function instanceOfCompanyType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CompanyTypeFromJSON(json: any): CompanyType {
    return CompanyTypeFromJSONTyped(json, false);
}

export function CompanyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompanyType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'companyName': !exists(json, 'companyName') ? undefined : json['companyName'],
        'companyName2': !exists(json, 'companyName2') ? undefined : json['companyName2'],
        'companyName3': !exists(json, 'companyName3') ? undefined : json['companyName3'],
        'alternateName': !exists(json, 'alternateName') ? undefined : json['alternateName'],
        'alternateLanguage': !exists(json, 'alternateLanguage') ? undefined : json['alternateLanguage'],
        'businessLocale': !exists(json, 'businessLocale') ? undefined : json['businessLocale'],
        'iATAInfo': !exists(json, 'iATAInfo') ? undefined : IATAInfoTypeFromJSON(json['iATAInfo']),
        'businessTitle': !exists(json, 'businessTitle') ? undefined : json['businessTitle'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'language': !exists(json, 'language') ? undefined : json['language'],
        'nationality': !exists(json, 'nationality') ? undefined : json['nationality'],
        'nationalityDescription': !exists(json, 'nationalityDescription') ? undefined : json['nationalityDescription'],
        'commissionCode': !exists(json, 'commissionCode') ? undefined : json['commissionCode'],
        'creditRating': !exists(json, 'creditRating') ? undefined : json['creditRating'],
        'corporateIdType': !exists(json, 'corporateIdType') ? undefined : json['corporateIdType'],
        'vipStatus': !exists(json, 'vipStatus') ? undefined : json['vipStatus'],
        'vipDescription': !exists(json, 'vipDescription') ? undefined : json['vipDescription'],
        'blacklist': !exists(json, 'blacklist') ? undefined : json['blacklist'],
    };
}

export function CompanyTypeToJSON(value?: CompanyType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'companyName': value.companyName,
        'companyName2': value.companyName2,
        'companyName3': value.companyName3,
        'alternateName': value.alternateName,
        'alternateLanguage': value.alternateLanguage,
        'businessLocale': value.businessLocale,
        'iATAInfo': IATAInfoTypeToJSON(value.iATAInfo),
        'businessTitle': value.businessTitle,
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'language': value.language,
        'nationality': value.nationality,
        'nationalityDescription': value.nationalityDescription,
        'commissionCode': value.commissionCode,
        'creditRating': value.creditRating,
        'corporateIdType': value.corporateIdType,
        'vipStatus': value.vipStatus,
        'vipDescription': value.vipDescription,
        'blacklist': value.blacklist,
    };
}

