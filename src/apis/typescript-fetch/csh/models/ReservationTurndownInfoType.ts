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
import type { TurndownStatusType } from './TurndownStatusType';
import {
    TurndownStatusTypeFromJSON,
    TurndownStatusTypeFromJSONTyped,
    TurndownStatusTypeToJSON,
} from './TurndownStatusType';

/**
 * Turndown information for a reservation
 * @export
 * @interface ReservationTurndownInfoType
 */
export interface ReservationTurndownInfoType {
    /**
     * Indicates if turndown is allowed or not
     * @type {boolean}
     * @memberof ReservationTurndownInfoType
     */
    allowed?: boolean;
    /**
     * 
     * @type {TurndownStatusType}
     * @memberof ReservationTurndownInfoType
     */
    status?: TurndownStatusType;
}

/**
 * Check if a given object implements the ReservationTurndownInfoType interface.
 */
export function instanceOfReservationTurndownInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationTurndownInfoTypeFromJSON(json: any): ReservationTurndownInfoType {
    return ReservationTurndownInfoTypeFromJSONTyped(json, false);
}

export function ReservationTurndownInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationTurndownInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'allowed': !exists(json, 'allowed') ? undefined : json['allowed'],
        'status': !exists(json, 'status') ? undefined : TurndownStatusTypeFromJSON(json['status']),
    };
}

export function ReservationTurndownInfoTypeToJSON(value?: ReservationTurndownInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'allowed': value.allowed,
        'status': TurndownStatusTypeToJSON(value.status),
    };
}

