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
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { UpsellInfoType } from './UpsellInfoType';
import {
    UpsellInfoTypeFromJSON,
    UpsellInfoTypeFromJSONTyped,
    UpsellInfoTypeToJSON,
} from './UpsellInfoType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response object to the that contains information for a reservation prior to upgrade.
 * @export
 * @interface ReservationUpsellInfo
 */
export interface ReservationUpsellInfo {
    /**
     * 
     * @type {UpsellInfoType}
     * @memberof ReservationUpsellInfo
     */
    upsellInfo?: UpsellInfoType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ReservationUpsellInfo
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReservationUpsellInfo
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ReservationUpsellInfo interface.
 */
export function instanceOfReservationUpsellInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationUpsellInfoFromJSON(json: any): ReservationUpsellInfo {
    return ReservationUpsellInfoFromJSONTyped(json, false);
}

export function ReservationUpsellInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationUpsellInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'upsellInfo': !exists(json, 'upsellInfo') ? undefined : UpsellInfoTypeFromJSON(json['upsellInfo']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ReservationUpsellInfoToJSON(value?: ReservationUpsellInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'upsellInfo': UpsellInfoTypeToJSON(value.upsellInfo),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

