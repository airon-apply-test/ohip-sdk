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
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface ReinstateCheckOutRequest
 */
export interface ReinstateCheckOutRequest {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof ReinstateCheckOutRequest
     */
    hotelId?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof ReinstateCheckOutRequest
     */
    reservationId?: ReservationId;
    /**
     * Indicates whether interfaces should be notified.
     * @type {boolean}
     * @memberof ReinstateCheckOutRequest
     */
    notifyInterfaces?: boolean;
    /**
     * Indicate whether user wants to reinstate even if some other guest already checked in into same room.
     * @type {boolean}
     * @memberof ReinstateCheckOutRequest
     */
    overrideFlag?: boolean;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof ReinstateCheckOutRequest
     */
    cashierId?: number;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ReinstateCheckOutRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReinstateCheckOutRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ReinstateCheckOutRequest interface.
 */
export function instanceOfReinstateCheckOutRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReinstateCheckOutRequestFromJSON(json: any): ReinstateCheckOutRequest {
    return ReinstateCheckOutRequestFromJSONTyped(json, false);
}

export function ReinstateCheckOutRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReinstateCheckOutRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'notifyInterfaces': !exists(json, 'notifyInterfaces') ? undefined : json['notifyInterfaces'],
        'overrideFlag': !exists(json, 'overrideFlag') ? undefined : json['overrideFlag'],
        'cashierId': !exists(json, 'cashierId') ? undefined : json['cashierId'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ReinstateCheckOutRequestToJSON(value?: ReinstateCheckOutRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'notifyInterfaces': value.notifyInterfaces,
        'overrideFlag': value.overrideFlag,
        'cashierId': value.cashierId,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

