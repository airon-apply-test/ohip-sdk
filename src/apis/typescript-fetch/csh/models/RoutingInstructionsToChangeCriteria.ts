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
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';
import type { RoutingInfoType } from './RoutingInfoType';
import {
    RoutingInfoTypeFromJSON,
    RoutingInfoTypeFromJSONTyped,
    RoutingInfoTypeToJSON,
} from './RoutingInfoType';
import type { RoutingInstructionsToChangeCriteriaComp } from './RoutingInstructionsToChangeCriteriaComp';
import {
    RoutingInstructionsToChangeCriteriaCompFromJSON,
    RoutingInstructionsToChangeCriteriaCompFromJSONTyped,
    RoutingInstructionsToChangeCriteriaCompToJSON,
} from './RoutingInstructionsToChangeCriteriaComp';
import type { RoutingInstructionsToChangeCriteriaFolio } from './RoutingInstructionsToChangeCriteriaFolio';
import {
    RoutingInstructionsToChangeCriteriaFolioFromJSON,
    RoutingInstructionsToChangeCriteriaFolioFromJSONTyped,
    RoutingInstructionsToChangeCriteriaFolioToJSON,
} from './RoutingInstructionsToChangeCriteriaFolio';
import type { RoutingInstructionsToChangeCriteriaRequest } from './RoutingInstructionsToChangeCriteriaRequest';
import {
    RoutingInstructionsToChangeCriteriaRequestFromJSON,
    RoutingInstructionsToChangeCriteriaRequestFromJSONTyped,
    RoutingInstructionsToChangeCriteriaRequestToJSON,
} from './RoutingInstructionsToChangeCriteriaRequest';
import type { RoutingInstructionsToChangeCriteriaRoom } from './RoutingInstructionsToChangeCriteriaRoom';
import {
    RoutingInstructionsToChangeCriteriaRoomFromJSON,
    RoutingInstructionsToChangeCriteriaRoomFromJSONTyped,
    RoutingInstructionsToChangeCriteriaRoomToJSON,
} from './RoutingInstructionsToChangeCriteriaRoom';

/**
 * Transactions and scheduled instructions included in this routing element will be replaced with the new element.
 * @export
 * @interface RoutingInstructionsToChangeCriteria
 */
export interface RoutingInstructionsToChangeCriteria {
    /**
     * 
     * @type {RoutingInstructionsToChangeCriteriaFolio}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    folio?: RoutingInstructionsToChangeCriteriaFolio;
    /**
     * 
     * @type {RoutingInstructionsToChangeCriteriaRoom}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    room?: RoutingInstructionsToChangeCriteriaRoom;
    /**
     * 
     * @type {RoutingInstructionsToChangeCriteriaComp}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    comp?: RoutingInstructionsToChangeCriteriaComp;
    /**
     * 
     * @type {RoutingInstructionsToChangeCriteriaRequest}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    request?: RoutingInstructionsToChangeCriteriaRequest;
    /**
     * Hotel context of the reservation.
     * @type {string}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    hotelId?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    reservationId?: ReservationId;
    /**
     * On a successful update, the transactions that are already posted in the guest's folio will be re-organized based on the configured instructions.
     * @type {boolean}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    refreshFolio?: boolean;
    /**
     * This flag indicates if postings that can be refreshed need to be part of the response when a routing instruction is created, updated or deleted.
     * @type {boolean}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    retrievePostingsForRoomRouting?: boolean;
    /**
     * 
     * @type {RoutingInfoType}
     * @memberof RoutingInstructionsToChangeCriteria
     */
    newRoutingInfo?: RoutingInfoType;
}

/**
 * Check if a given object implements the RoutingInstructionsToChangeCriteria interface.
 */
export function instanceOfRoutingInstructionsToChangeCriteria(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoutingInstructionsToChangeCriteriaFromJSON(json: any): RoutingInstructionsToChangeCriteria {
    return RoutingInstructionsToChangeCriteriaFromJSONTyped(json, false);
}

export function RoutingInstructionsToChangeCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInstructionsToChangeCriteria {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'folio': !exists(json, 'folio') ? undefined : RoutingInstructionsToChangeCriteriaFolioFromJSON(json['folio']),
        'room': !exists(json, 'room') ? undefined : RoutingInstructionsToChangeCriteriaRoomFromJSON(json['room']),
        'comp': !exists(json, 'comp') ? undefined : RoutingInstructionsToChangeCriteriaCompFromJSON(json['comp']),
        'request': !exists(json, 'request') ? undefined : RoutingInstructionsToChangeCriteriaRequestFromJSON(json['request']),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'refreshFolio': !exists(json, 'refreshFolio') ? undefined : json['refreshFolio'],
        'retrievePostingsForRoomRouting': !exists(json, 'retrievePostingsForRoomRouting') ? undefined : json['retrievePostingsForRoomRouting'],
        'newRoutingInfo': !exists(json, 'newRoutingInfo') ? undefined : RoutingInfoTypeFromJSON(json['newRoutingInfo']),
    };
}

export function RoutingInstructionsToChangeCriteriaToJSON(value?: RoutingInstructionsToChangeCriteria | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'folio': RoutingInstructionsToChangeCriteriaFolioToJSON(value.folio),
        'room': RoutingInstructionsToChangeCriteriaRoomToJSON(value.room),
        'comp': RoutingInstructionsToChangeCriteriaCompToJSON(value.comp),
        'request': RoutingInstructionsToChangeCriteriaRequestToJSON(value.request),
        'hotelId': value.hotelId,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'refreshFolio': value.refreshFolio,
        'retrievePostingsForRoomRouting': value.retrievePostingsForRoomRouting,
        'newRoutingInfo': RoutingInfoTypeToJSON(value.newRoutingInfo),
    };
}

