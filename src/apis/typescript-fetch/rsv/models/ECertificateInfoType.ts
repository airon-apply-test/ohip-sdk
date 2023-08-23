/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { ECertificateInfoTypeHotels } from './ECertificateInfoTypeHotels';
import {
    ECertificateInfoTypeHotelsFromJSON,
    ECertificateInfoTypeHotelsFromJSONTyped,
    ECertificateInfoTypeHotelsToJSON,
} from './ECertificateInfoTypeHotels';

/**
 * E-Certificates details.
 * @export
 * @interface ECertificateInfoType
 */
export interface ECertificateInfoType {
    /**
     * User defined certificate code.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    certificateType?: string;
    /**
     * Membership type to which the certificate is linked to.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    membershipType?: string;
    /**
     * Award type to which the certificate is linked to.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    awardCode?: string;
    /**
     * Promotion code to which certificate is attached to.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    promotionCode?: string;
    /**
     * Voucher benefit code attached to the certificate.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    voucherBenefitCode?: string;
    /**
     * 
     * @type {ECertificateInfoTypeHotels}
     * @memberof ECertificateInfoType
     */
    hotels?: ECertificateInfoTypeHotels;
    /**
     * Description about the certificate.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    description?: string;
    /**
     * Detail description about the certificate.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    longDescription?: string;
    /**
     * Label for the certificate.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    label?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof ECertificateInfoType
     */
    value?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof ECertificateInfoType
     */
    cost?: CurrencyAmountType;
    /**
     * Summary of Benefits attached to this ECertificate.
     * @type {string}
     * @memberof ECertificateInfoType
     */
    benefitSummary?: string;
}

/**
 * Check if a given object implements the ECertificateInfoType interface.
 */
export function instanceOfECertificateInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ECertificateInfoTypeFromJSON(json: any): ECertificateInfoType {
    return ECertificateInfoTypeFromJSONTyped(json, false);
}

export function ECertificateInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'certificateType': !exists(json, 'certificateType') ? undefined : json['certificateType'],
        'membershipType': !exists(json, 'membershipType') ? undefined : json['membershipType'],
        'awardCode': !exists(json, 'awardCode') ? undefined : json['awardCode'],
        'promotionCode': !exists(json, 'promotionCode') ? undefined : json['promotionCode'],
        'voucherBenefitCode': !exists(json, 'voucherBenefitCode') ? undefined : json['voucherBenefitCode'],
        'hotels': !exists(json, 'hotels') ? undefined : ECertificateInfoTypeHotelsFromJSON(json['hotels']),
        'description': !exists(json, 'description') ? undefined : json['description'],
        'longDescription': !exists(json, 'longDescription') ? undefined : json['longDescription'],
        'label': !exists(json, 'label') ? undefined : json['label'],
        'value': !exists(json, 'value') ? undefined : CurrencyAmountTypeFromJSON(json['value']),
        'cost': !exists(json, 'cost') ? undefined : CurrencyAmountTypeFromJSON(json['cost']),
        'benefitSummary': !exists(json, 'benefitSummary') ? undefined : json['benefitSummary'],
    };
}

export function ECertificateInfoTypeToJSON(value?: ECertificateInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'certificateType': value.certificateType,
        'membershipType': value.membershipType,
        'awardCode': value.awardCode,
        'promotionCode': value.promotionCode,
        'voucherBenefitCode': value.voucherBenefitCode,
        'hotels': ECertificateInfoTypeHotelsToJSON(value.hotels),
        'description': value.description,
        'longDescription': value.longDescription,
        'label': value.label,
        'value': CurrencyAmountTypeToJSON(value.value),
        'cost': CurrencyAmountTypeToJSON(value.cost),
        'benefitSummary': value.benefitSummary,
    };
}

