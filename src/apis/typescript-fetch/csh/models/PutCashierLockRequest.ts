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
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PutCashierLockRequest
 */
export interface PutCashierLockRequest {
    /**
     * The number of seconds to add to the lock's validity time.
     * @type {number}
     * @memberof PutCashierLockRequest
     */
    additionalTimeToLive?: number;
    /**
     * The Cashier Lock Handle to pass along with operation which required cashier to be locked.
     * @type {number}
     * @memberof PutCashierLockRequest
     */
    cashierLockHandle?: number;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PutCashierLockRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutCashierLockRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PutCashierLockRequest interface.
 */
export function instanceOfPutCashierLockRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PutCashierLockRequestFromJSON(json: any): PutCashierLockRequest {
    return PutCashierLockRequestFromJSONTyped(json, false);
}

export function PutCashierLockRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutCashierLockRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'additionalTimeToLive': !exists(json, 'additionalTimeToLive') ? undefined : json['additionalTimeToLive'],
        'cashierLockHandle': !exists(json, 'cashierLockHandle') ? undefined : json['cashierLockHandle'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PutCashierLockRequestToJSON(value?: PutCashierLockRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'additionalTimeToLive': value.additionalTimeToLive,
        'cashierLockHandle': value.cashierLockHandle,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

