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
import type { ChargeCriteriaType } from './ChargeCriteriaType';
import {
    ChargeCriteriaTypeFromJSON,
    ChargeCriteriaTypeFromJSONTyped,
    ChargeCriteriaTypeToJSON,
} from './ChargeCriteriaType';
import type { CheckPostBillingChargesOptionsTypeInnerRoomRouting } from './CheckPostBillingChargesOptionsTypeInnerRoomRouting';
import {
    CheckPostBillingChargesOptionsTypeInnerRoomRoutingFromJSON,
    CheckPostBillingChargesOptionsTypeInnerRoomRoutingFromJSONTyped,
    CheckPostBillingChargesOptionsTypeInnerRoomRoutingToJSON,
} from './CheckPostBillingChargesOptionsTypeInnerRoomRouting';

/**
 * 
 * @export
 * @interface CheckPostBillingChargesOptionsTypeInner
 */
export interface CheckPostBillingChargesOptionsTypeInner {
    /**
     * 
     * @type {ChargeCriteriaType}
     * @memberof CheckPostBillingChargesOptionsTypeInner
     */
    consumable?: ChargeCriteriaType;
    /**
     * Determines if the package allowance is available to be consumed.
     * @type {boolean}
     * @memberof CheckPostBillingChargesOptionsTypeInner
     */
    packageAllowance?: boolean;
    /**
     * 
     * @type {CheckPostBillingChargesOptionsTypeInnerRoomRouting}
     * @memberof CheckPostBillingChargesOptionsTypeInner
     */
    roomRouting?: CheckPostBillingChargesOptionsTypeInnerRoomRouting;
}

/**
 * Check if a given object implements the CheckPostBillingChargesOptionsTypeInner interface.
 */
export function instanceOfCheckPostBillingChargesOptionsTypeInner(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CheckPostBillingChargesOptionsTypeInnerFromJSON(json: any): CheckPostBillingChargesOptionsTypeInner {
    return CheckPostBillingChargesOptionsTypeInnerFromJSONTyped(json, false);
}

export function CheckPostBillingChargesOptionsTypeInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): CheckPostBillingChargesOptionsTypeInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'consumable': !exists(json, 'consumable') ? undefined : ChargeCriteriaTypeFromJSON(json['consumable']),
        'packageAllowance': !exists(json, 'packageAllowance') ? undefined : json['packageAllowance'],
        'roomRouting': !exists(json, 'roomRouting') ? undefined : CheckPostBillingChargesOptionsTypeInnerRoomRoutingFromJSON(json['roomRouting']),
    };
}

export function CheckPostBillingChargesOptionsTypeInnerToJSON(value?: CheckPostBillingChargesOptionsTypeInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'consumable': ChargeCriteriaTypeToJSON(value.consumable),
        'packageAllowance': value.packageAllowance,
        'roomRouting': CheckPostBillingChargesOptionsTypeInnerRoomRoutingToJSON(value.roomRouting),
    };
}

