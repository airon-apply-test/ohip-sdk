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
import type { FolioWindowType } from './FolioWindowType';
import {
    FolioWindowTypeFromJSON,
    FolioWindowTypeFromJSONTyped,
    FolioWindowTypeToJSON,
} from './FolioWindowType';
import type { ReservationInfoType } from './ReservationInfoType';
import {
    ReservationInfoTypeFromJSON,
    ReservationInfoTypeFromJSONTyped,
    ReservationInfoTypeToJSON,
} from './ReservationInfoType';

/**
 * Each folio will include summary and/or detailed information.
 * @export
 * @interface ReservationFolioInfoType
 */
export interface ReservationFolioInfoType {
    /**
     * 
     * @type {ReservationInfoType}
     * @memberof ReservationFolioInfoType
     */
    reservationInfo?: ReservationInfoType;
    /**
     * The Folio Window information of the reservation.
     * @type {Array<FolioWindowType>}
     * @memberof ReservationFolioInfoType
     */
    folioWindows?: Array<FolioWindowType>;
    /**
     * Folio History element contains all generated folios
     * @type {Array<FolioWindowType>}
     * @memberof ReservationFolioInfoType
     */
    folioHistory?: Array<FolioWindowType>;
    /**
     * Computed flag specifying that the reservation has met the criteria which allows charges to be posted after being checked-out.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    postStayChargeAllowed?: boolean;
    /**
     * Computed flag specifying that the reservation has met the criteria which allows charges to be posted before checking-in.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    preStayChargeAllowed?: boolean;
    /**
     * Flag specifying that the reservation can be auto check in when Auto Checkin Pseudo Room parameter is active.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    autoCheckInAllowed?: boolean;
    /**
     * Flag specifying that the reservation can post if the reservation status was No Show or Cancelled.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    postToNoShowCancelAllowed?: boolean;
    /**
     * Stamp Duty posted in at least one folio window.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    stampDutyExists?: boolean;
    /**
     * Flag to check if the room and tax are already posted.
     * @type {boolean}
     * @memberof ReservationFolioInfoType
     */
    roomAndTaxPosted?: boolean;
    /**
     * Flag applicable only when ALLOW_DEFERRED_TAXES is ON for the resort. Set to true only when there are any unsettled transactions on any of the windows with the deferred tax entry posted. Set to false only when there are unsettled transactions on any of the windows without deferred tax entry.
     * @type {string}
     * @memberof ReservationFolioInfoType
     */
    deferredTaxesPosted?: string;
}

/**
 * Check if a given object implements the ReservationFolioInfoType interface.
 */
export function instanceOfReservationFolioInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationFolioInfoTypeFromJSON(json: any): ReservationFolioInfoType {
    return ReservationFolioInfoTypeFromJSONTyped(json, false);
}

export function ReservationFolioInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationFolioInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationInfo': !exists(json, 'reservationInfo') ? undefined : ReservationInfoTypeFromJSON(json['reservationInfo']),
        'folioWindows': !exists(json, 'folioWindows') ? undefined : ((json['folioWindows'] as Array<any>).map(FolioWindowTypeFromJSON)),
        'folioHistory': !exists(json, 'folioHistory') ? undefined : ((json['folioHistory'] as Array<any>).map(FolioWindowTypeFromJSON)),
        'postStayChargeAllowed': !exists(json, 'postStayChargeAllowed') ? undefined : json['postStayChargeAllowed'],
        'preStayChargeAllowed': !exists(json, 'preStayChargeAllowed') ? undefined : json['preStayChargeAllowed'],
        'autoCheckInAllowed': !exists(json, 'autoCheckInAllowed') ? undefined : json['autoCheckInAllowed'],
        'postToNoShowCancelAllowed': !exists(json, 'postToNoShowCancelAllowed') ? undefined : json['postToNoShowCancelAllowed'],
        'stampDutyExists': !exists(json, 'stampDutyExists') ? undefined : json['stampDutyExists'],
        'roomAndTaxPosted': !exists(json, 'roomAndTaxPosted') ? undefined : json['roomAndTaxPosted'],
        'deferredTaxesPosted': !exists(json, 'deferredTaxesPosted') ? undefined : json['deferredTaxesPosted'],
    };
}

export function ReservationFolioInfoTypeToJSON(value?: ReservationFolioInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationInfo': ReservationInfoTypeToJSON(value.reservationInfo),
        'folioWindows': value.folioWindows === undefined ? undefined : ((value.folioWindows as Array<any>).map(FolioWindowTypeToJSON)),
        'folioHistory': value.folioHistory === undefined ? undefined : ((value.folioHistory as Array<any>).map(FolioWindowTypeToJSON)),
        'postStayChargeAllowed': value.postStayChargeAllowed,
        'preStayChargeAllowed': value.preStayChargeAllowed,
        'autoCheckInAllowed': value.autoCheckInAllowed,
        'postToNoShowCancelAllowed': value.postToNoShowCancelAllowed,
        'stampDutyExists': value.stampDutyExists,
        'roomAndTaxPosted': value.roomAndTaxPosted,
        'deferredTaxesPosted': value.deferredTaxesPosted,
    };
}

