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
import type { ProfileRoutingInstructionsType } from './ProfileRoutingInstructionsType';
import {
    ProfileRoutingInstructionsTypeFromJSON,
    ProfileRoutingInstructionsTypeFromJSONTyped,
    ProfileRoutingInstructionsTypeToJSON,
} from './ProfileRoutingInstructionsType';

/**
 * The type contains routing instructions for the profile.
 * @export
 * @interface ProfileCashieringDetailType
 */
export interface ProfileCashieringDetailType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof ProfileCashieringDetailType
     */
    paymentMethod?: CodeDescriptionType;
    /**
     * 
     * @type {ProfileRoutingInstructionsType}
     * @memberof ProfileCashieringDetailType
     */
    routingInstructions?: ProfileRoutingInstructionsType;
    /**
     * Tax type code.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    taxType?: string;
    /**
     * Guest type code.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    fiscalGuestType?: string;
    /**
     * Hotel Code for which the routing instructions are provided for a profile.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the ProfileCashieringDetailType interface.
 */
export function instanceOfProfileCashieringDetailType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileCashieringDetailTypeFromJSON(json: any): ProfileCashieringDetailType {
    return ProfileCashieringDetailTypeFromJSONTyped(json, false);
}

export function ProfileCashieringDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCashieringDetailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'paymentMethod': !exists(json, 'paymentMethod') ? undefined : CodeDescriptionTypeFromJSON(json['paymentMethod']),
        'routingInstructions': !exists(json, 'routingInstructions') ? undefined : ProfileRoutingInstructionsTypeFromJSON(json['routingInstructions']),
        'taxType': !exists(json, 'taxType') ? undefined : json['taxType'],
        'fiscalGuestType': !exists(json, 'fiscalGuestType') ? undefined : json['fiscalGuestType'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function ProfileCashieringDetailTypeToJSON(value?: ProfileCashieringDetailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'paymentMethod': CodeDescriptionTypeToJSON(value.paymentMethod),
        'routingInstructions': ProfileRoutingInstructionsTypeToJSON(value.routingInstructions),
        'taxType': value.taxType,
        'fiscalGuestType': value.fiscalGuestType,
        'hotelId': value.hotelId,
    };
}

