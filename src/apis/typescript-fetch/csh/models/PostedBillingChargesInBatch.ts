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
import type { ReservationChargesInBatchInfoType } from './ReservationChargesInBatchInfoType';
import {
    ReservationChargesInBatchInfoTypeFromJSON,
    ReservationChargesInBatchInfoTypeFromJSONTyped,
    ReservationChargesInBatchInfoTypeToJSON,
} from './ReservationChargesInBatchInfoType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Operation response to post billing charge with list of reservations.
 * @export
 * @interface PostedBillingChargesInBatch
 */
export interface PostedBillingChargesInBatch {
    /**
     * Information regarding charges in batch result for each reservation.
     * @type {Array<ReservationChargesInBatchInfoType>}
     * @memberof PostedBillingChargesInBatch
     */
    reservationBatchChargeInfoList?: Array<ReservationChargesInBatchInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostedBillingChargesInBatch
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostedBillingChargesInBatch
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostedBillingChargesInBatch interface.
 */
export function instanceOfPostedBillingChargesInBatch(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostedBillingChargesInBatchFromJSON(json: any): PostedBillingChargesInBatch {
    return PostedBillingChargesInBatchFromJSONTyped(json, false);
}

export function PostedBillingChargesInBatchFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostedBillingChargesInBatch {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationBatchChargeInfoList': !exists(json, 'reservationBatchChargeInfoList') ? undefined : ((json['reservationBatchChargeInfoList'] as Array<any>).map(ReservationChargesInBatchInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostedBillingChargesInBatchToJSON(value?: PostedBillingChargesInBatch | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationBatchChargeInfoList': value.reservationBatchChargeInfoList === undefined ? undefined : ((value.reservationBatchChargeInfoList as Array<any>).map(ReservationChargesInBatchInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

