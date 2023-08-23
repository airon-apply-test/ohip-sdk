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
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { CustomChargeExemptionDateType } from './CustomChargeExemptionDateType';
import {
    CustomChargeExemptionDateTypeFromJSON,
    CustomChargeExemptionDateTypeFromJSONTyped,
    CustomChargeExemptionDateTypeToJSON,
} from './CustomChargeExemptionDateType';
import type { CustomChargeQuantityType } from './CustomChargeQuantityType';
import {
    CustomChargeQuantityTypeFromJSON,
    CustomChargeQuantityTypeFromJSONTyped,
    CustomChargeQuantityTypeToJSON,
} from './CustomChargeQuantityType';
import type { ExcludedDateType } from './ExcludedDateType';
import {
    ExcludedDateTypeFromJSON,
    ExcludedDateTypeFromJSONTyped,
    ExcludedDateTypeToJSON,
} from './ExcludedDateType';

/**
 * Contains custom charges exemption information.
 * @export
 * @interface CustomChargeExemptionType
 */
export interface CustomChargeExemptionType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof CustomChargeExemptionType
     */
    customChargesExemption?: CodeDescriptionType;
    /**
     * 
     * @type {CustomChargeQuantityType}
     * @memberof CustomChargeExemptionType
     */
    customChargeQuantity?: CustomChargeQuantityType;
    /**
     * Contains List of Custom Charge Exemption information for a day.
     * @type {Array<CustomChargeExemptionDateType>}
     * @memberof CustomChargeExemptionType
     */
    customChargeDates?: Array<CustomChargeExemptionDateType>;
    /**
     * Contains list of dates which are valid for custom charge exemptions.
     * @type {Array<ExcludedDateType>}
     * @memberof CustomChargeExemptionType
     */
    excludedDates?: Array<ExcludedDateType>;
    /**
     * Exemption Percentage
     * @type {number}
     * @memberof CustomChargeExemptionType
     */
    percentage?: number;
    /**
     * Flag specifying if custom charge exemptions is property level or not.
     * @type {boolean}
     * @memberof CustomChargeExemptionType
     */
    propertyExemption?: boolean;
}

/**
 * Check if a given object implements the CustomChargeExemptionType interface.
 */
export function instanceOfCustomChargeExemptionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CustomChargeExemptionTypeFromJSON(json: any): CustomChargeExemptionType {
    return CustomChargeExemptionTypeFromJSONTyped(json, false);
}

export function CustomChargeExemptionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CustomChargeExemptionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customChargesExemption': !exists(json, 'customChargesExemption') ? undefined : CodeDescriptionTypeFromJSON(json['customChargesExemption']),
        'customChargeQuantity': !exists(json, 'customChargeQuantity') ? undefined : CustomChargeQuantityTypeFromJSON(json['customChargeQuantity']),
        'customChargeDates': !exists(json, 'customChargeDates') ? undefined : ((json['customChargeDates'] as Array<any>).map(CustomChargeExemptionDateTypeFromJSON)),
        'excludedDates': !exists(json, 'excludedDates') ? undefined : ((json['excludedDates'] as Array<any>).map(ExcludedDateTypeFromJSON)),
        'percentage': !exists(json, 'percentage') ? undefined : json['percentage'],
        'propertyExemption': !exists(json, 'propertyExemption') ? undefined : json['propertyExemption'],
    };
}

export function CustomChargeExemptionTypeToJSON(value?: CustomChargeExemptionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'customChargesExemption': CodeDescriptionTypeToJSON(value.customChargesExemption),
        'customChargeQuantity': CustomChargeQuantityTypeToJSON(value.customChargeQuantity),
        'customChargeDates': value.customChargeDates === undefined ? undefined : ((value.customChargeDates as Array<any>).map(CustomChargeExemptionDateTypeToJSON)),
        'excludedDates': value.excludedDates === undefined ? undefined : ((value.excludedDates as Array<any>).map(ExcludedDateTypeToJSON)),
        'percentage': value.percentage,
        'propertyExemption': value.propertyExemption,
    };
}

