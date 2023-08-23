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
import type { PostDepositToGuestFolioType } from './PostDepositToGuestFolioType';
import {
    PostDepositToGuestFolioTypeFromJSON,
    PostDepositToGuestFolioTypeFromJSONTyped,
    PostDepositToGuestFolioTypeToJSON,
} from './PostDepositToGuestFolioType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface TransitionDepositToGuestFolioRequest
 */
export interface TransitionDepositToGuestFolioRequest {
    /**
     * 
     * @type {PostDepositToGuestFolioType}
     * @memberof TransitionDepositToGuestFolioRequest
     */
    criteria?: PostDepositToGuestFolioType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof TransitionDepositToGuestFolioRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TransitionDepositToGuestFolioRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the TransitionDepositToGuestFolioRequest interface.
 */
export function instanceOfTransitionDepositToGuestFolioRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TransitionDepositToGuestFolioRequestFromJSON(json: any): TransitionDepositToGuestFolioRequest {
    return TransitionDepositToGuestFolioRequestFromJSONTyped(json, false);
}

export function TransitionDepositToGuestFolioRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransitionDepositToGuestFolioRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : PostDepositToGuestFolioTypeFromJSON(json['criteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function TransitionDepositToGuestFolioRequestToJSON(value?: TransitionDepositToGuestFolioRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': PostDepositToGuestFolioTypeToJSON(value.criteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

