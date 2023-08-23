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
/**
 * Indicates the Credit Card EFT payment method authorization rules setup.
 * @export
 * @interface CCAuthorizationInstructionTypeAuthorizationSetup
 */
export interface CCAuthorizationInstructionTypeAuthorizationSetup {
    /**
     * Indicates if authorization at Check In is allowed for this payment method. False will indicate that no authorization will be done at Check In for this payment method.
     * @type {boolean}
     * @memberof CCAuthorizationInstructionTypeAuthorizationSetup
     */
    checkIn?: boolean;
    /**
     * Indicates if authorization during the stay of the guest is allowed for this payment method. False will indicate that no authorization will be done during the stay and prior to settlements for this payment method. The settlement request will be sent without an authorization request.
     * @type {boolean}
     * @memberof CCAuthorizationInstructionTypeAuthorizationSetup
     */
    stay?: boolean;
    /**
     * Indicates if authorization at Deposit is allowed for this payment method. False will indicate that no authorization will be done at the time of deposit payments, for this payment method. The settlement request will be sent without an authorization request.
     * @type {boolean}
     * @memberof CCAuthorizationInstructionTypeAuthorizationSetup
     */
    deposit?: boolean;
    /**
     * Indicates if this payment method is setup as a PayOnly, which does not require authorization to be done prior to settlement. The special settlement handling will take care of both Authorization and Settlement together.
     * @type {boolean}
     * @memberof CCAuthorizationInstructionTypeAuthorizationSetup
     */
    payOnly?: boolean;
}

/**
 * Check if a given object implements the CCAuthorizationInstructionTypeAuthorizationSetup interface.
 */
export function instanceOfCCAuthorizationInstructionTypeAuthorizationSetup(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CCAuthorizationInstructionTypeAuthorizationSetupFromJSON(json: any): CCAuthorizationInstructionTypeAuthorizationSetup {
    return CCAuthorizationInstructionTypeAuthorizationSetupFromJSONTyped(json, false);
}

export function CCAuthorizationInstructionTypeAuthorizationSetupFromJSONTyped(json: any, ignoreDiscriminator: boolean): CCAuthorizationInstructionTypeAuthorizationSetup {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkIn': !exists(json, 'checkIn') ? undefined : json['checkIn'],
        'stay': !exists(json, 'stay') ? undefined : json['stay'],
        'deposit': !exists(json, 'deposit') ? undefined : json['deposit'],
        'payOnly': !exists(json, 'payOnly') ? undefined : json['payOnly'],
    };
}

export function CCAuthorizationInstructionTypeAuthorizationSetupToJSON(value?: CCAuthorizationInstructionTypeAuthorizationSetup | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkIn': value.checkIn,
        'stay': value.stay,
        'deposit': value.deposit,
        'payOnly': value.payOnly,
    };
}

