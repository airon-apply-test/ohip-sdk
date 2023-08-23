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
import type { ProfileDeliveryModuleType } from './ProfileDeliveryModuleType';
import {
    ProfileDeliveryModuleTypeFromJSON,
    ProfileDeliveryModuleTypeFromJSONTyped,
    ProfileDeliveryModuleTypeToJSON,
} from './ProfileDeliveryModuleType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Delivery Information type to the profile.
 * @export
 * @interface ProfileDeliveryMethod
 */
export interface ProfileDeliveryMethod {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ProfileDeliveryMethod
     */
    deliveryId?: UniqueIDType;
    /**
     * Delivery type can have a value EMAIL, ELECTRONIC etc and it depends on the parameter set in OPERA Control.
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    deliveryType?: string;
    /**
     * Delivery value holds the corresponding value of the delivery type..
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    deliveryValue?: string;
    /**
     * Property that has delivery methods configured.
     * @type {string}
     * @memberof ProfileDeliveryMethod
     */
    hotelId?: string;
    /**
     * 
     * @type {ProfileDeliveryModuleType}
     * @memberof ProfileDeliveryMethod
     */
    deliveryModule?: ProfileDeliveryModuleType;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof ProfileDeliveryMethod
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof ProfileDeliveryMethod
     */
    orderSequence?: number;
}

/**
 * Check if a given object implements the ProfileDeliveryMethod interface.
 */
export function instanceOfProfileDeliveryMethod(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileDeliveryMethodFromJSON(json: any): ProfileDeliveryMethod {
    return ProfileDeliveryMethodFromJSONTyped(json, false);
}

export function ProfileDeliveryMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileDeliveryMethod {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deliveryId': !exists(json, 'deliveryId') ? undefined : UniqueIDTypeFromJSON(json['deliveryId']),
        'deliveryType': !exists(json, 'deliveryType') ? undefined : json['deliveryType'],
        'deliveryValue': !exists(json, 'deliveryValue') ? undefined : json['deliveryValue'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'deliveryModule': !exists(json, 'deliveryModule') ? undefined : ProfileDeliveryModuleTypeFromJSON(json['deliveryModule']),
        'primaryInd': !exists(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !exists(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}

export function ProfileDeliveryMethodToJSON(value?: ProfileDeliveryMethod | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deliveryId': UniqueIDTypeToJSON(value.deliveryId),
        'deliveryType': value.deliveryType,
        'deliveryValue': value.deliveryValue,
        'hotelId': value.hotelId,
        'deliveryModule': ProfileDeliveryModuleTypeToJSON(value.deliveryModule),
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
    };
}

