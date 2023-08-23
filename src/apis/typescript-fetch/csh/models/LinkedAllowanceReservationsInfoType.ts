/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { LinkedAllowanceReservationInfoType } from './LinkedAllowanceReservationInfoType';
import {
    LinkedAllowanceReservationInfoTypeFromJSON,
    LinkedAllowanceReservationInfoTypeFromJSONTyped,
    LinkedAllowanceReservationInfoTypeToJSON,
} from './LinkedAllowanceReservationInfoType';

/**
 * List of reservation with linked package allowances.
 * @export
 * @interface LinkedAllowanceReservationsInfoType
 */
export interface LinkedAllowanceReservationsInfoType {
    /**
     * The resort code.
     * @type {string}
     * @memberof LinkedAllowanceReservationsInfoType
     */
    hotelId?: string;
    /**
     * List of Package Allowance Source reservations.
     * @type {Array<LinkedAllowanceReservationInfoType>}
     * @memberof LinkedAllowanceReservationsInfoType
     */
    sourceReservations?: Array<LinkedAllowanceReservationInfoType>;
    /**
     * 
     * @type {LinkedAllowanceReservationInfoType}
     * @memberof LinkedAllowanceReservationsInfoType
     */
    targetReservation?: LinkedAllowanceReservationInfoType;
}

/**
 * Check if a given object implements the LinkedAllowanceReservationsInfoType interface.
 */
export function instanceOfLinkedAllowanceReservationsInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function LinkedAllowanceReservationsInfoTypeFromJSON(json: any): LinkedAllowanceReservationsInfoType {
    return LinkedAllowanceReservationsInfoTypeFromJSONTyped(json, false);
}

export function LinkedAllowanceReservationsInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): LinkedAllowanceReservationsInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'sourceReservations': !exists(json, 'sourceReservations') ? undefined : ((json['sourceReservations'] as Array<any>).map(LinkedAllowanceReservationInfoTypeFromJSON)),
        'targetReservation': !exists(json, 'targetReservation') ? undefined : LinkedAllowanceReservationInfoTypeFromJSON(json['targetReservation']),
    };
}

export function LinkedAllowanceReservationsInfoTypeToJSON(value?: LinkedAllowanceReservationsInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'sourceReservations': value.sourceReservations === undefined ? undefined : ((value.sourceReservations as Array<any>).map(LinkedAllowanceReservationInfoTypeToJSON)),
        'targetReservation': LinkedAllowanceReservationInfoTypeToJSON(value.targetReservation),
    };
}

