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
import type { ECertificateConsumeSourceType } from './ECertificateConsumeSourceType';
import {
    ECertificateConsumeSourceTypeFromJSON,
    ECertificateConsumeSourceTypeFromJSONTyped,
    ECertificateConsumeSourceTypeToJSON,
} from './ECertificateConsumeSourceType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * E-Certificates details.
 * @export
 * @interface ECertificateConsumptionType
 */
export interface ECertificateConsumptionType {
    /**
     * Property where certificate was consumed for.
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    hotelId?: string;
    /**
     * 
     * @type {ECertificateConsumeSourceType}
     * @memberof ECertificateConsumptionType
     */
    source?: ECertificateConsumeSourceType;
    /**
     * Last name of the person who consumed the certificate..
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    surname?: string;
    /**
     * First name of the person who consumed the certificate.
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    firstName?: string;
    /**
     * Middle name of the person who consumed the certificate.
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    middleName?: string;
    /**
     * Email of the person who consumed the certificate.
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    email?: string;
    /**
     * Date the certificate was consumed.
     * @type {Date}
     * @memberof ECertificateConsumptionType
     */
    date?: Date;
    /**
     * Application user who created the consumption.
     * @type {string}
     * @memberof ECertificateConsumptionType
     */
    userName?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ECertificateConsumptionType
     */
    referenceId?: UniqueIDType;
}

/**
 * Check if a given object implements the ECertificateConsumptionType interface.
 */
export function instanceOfECertificateConsumptionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ECertificateConsumptionTypeFromJSON(json: any): ECertificateConsumptionType {
    return ECertificateConsumptionTypeFromJSONTyped(json, false);
}

export function ECertificateConsumptionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateConsumptionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'source': !exists(json, 'source') ? undefined : ECertificateConsumeSourceTypeFromJSON(json['source']),
        'surname': !exists(json, 'surname') ? undefined : json['surname'],
        'firstName': !exists(json, 'firstName') ? undefined : json['firstName'],
        'middleName': !exists(json, 'middleName') ? undefined : json['middleName'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'date': !exists(json, 'date') ? undefined : (new Date(json['date'])),
        'userName': !exists(json, 'userName') ? undefined : json['userName'],
        'referenceId': !exists(json, 'referenceId') ? undefined : UniqueIDTypeFromJSON(json['referenceId']),
    };
}

export function ECertificateConsumptionTypeToJSON(value?: ECertificateConsumptionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'source': ECertificateConsumeSourceTypeToJSON(value.source),
        'surname': value.surname,
        'firstName': value.firstName,
        'middleName': value.middleName,
        'email': value.email,
        'date': value.date === undefined ? undefined : (value.date.toISOString().substr(0,10)),
        'userName': value.userName,
        'referenceId': UniqueIDTypeToJSON(value.referenceId),
    };
}

