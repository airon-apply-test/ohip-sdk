/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
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
     * Alternate Name of the Company. Mainly, it's the name of the company written in the Alternate Language.
     * @type {string}
     * @memberof CompanyType
     */
    alternateName?: string;
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
     * Commission code of the company.
     * @type {string}
     * @memberof CompanyType
     */
    commissionCode?: string;
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
        'alternateName': !exists(json, 'alternateName') ? undefined : json['alternateName'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'language': !exists(json, 'language') ? undefined : json['language'],
        'nationality': !exists(json, 'nationality') ? undefined : json['nationality'],
        'commissionCode': !exists(json, 'commissionCode') ? undefined : json['commissionCode'],
        'vipStatus': !exists(json, 'vipStatus') ? undefined : json['vipStatus'],
        'vipDescription': !exists(json, 'vipDescription') ? undefined : json['vipDescription'],
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
        'alternateName': value.alternateName,
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'language': value.language,
        'nationality': value.nationality,
        'commissionCode': value.commissionCode,
        'vipStatus': value.vipStatus,
        'vipDescription': value.vipDescription,
    };
}

