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
import type { DateRangeType } from './DateRangeType';
import {
    DateRangeTypeFromJSON,
    DateRangeTypeFromJSONTyped,
    DateRangeTypeToJSON,
} from './DateRangeType';
import type { ErrorType } from './ErrorType';
import {
    ErrorTypeFromJSON,
    ErrorTypeFromJSONTyped,
    ErrorTypeToJSON,
} from './ErrorType';
import type { VaultHTTPTransactionMessageTypeAuthorizationApproval } from './VaultHTTPTransactionMessageTypeAuthorizationApproval';
import {
    VaultHTTPTransactionMessageTypeAuthorizationApprovalFromJSON,
    VaultHTTPTransactionMessageTypeAuthorizationApprovalFromJSONTyped,
    VaultHTTPTransactionMessageTypeAuthorizationApprovalToJSON,
} from './VaultHTTPTransactionMessageTypeAuthorizationApproval';
import type { VaultHTTPTransactionType } from './VaultHTTPTransactionType';
import {
    VaultHTTPTransactionTypeFromJSON,
    VaultHTTPTransactionTypeFromJSONTyped,
    VaultHTTPTransactionTypeToJSON,
} from './VaultHTTPTransactionType';

/**
 * 
 * @export
 * @interface VaultHTTPTransactionMessageType
 */
export interface VaultHTTPTransactionMessageType {
    /**
     * The hotel context of the transaction.
     * @type {string}
     * @memberof VaultHTTPTransactionMessageType
     */
    hotelId?: string;
    /**
     * The HTTP request entity content. The needs to use escape characters.
     * @type {string}
     * @memberof VaultHTTPTransactionMessageType
     */
    escapedRequestContent?: string;
    /**
     * The HTTP response entity content. The needs to use escape characters.
     * @type {string}
     * @memberof VaultHTTPTransactionMessageType
     */
    escapedResponseContent?: string;
    /**
     * 
     * @type {DateRangeType}
     * @memberof VaultHTTPTransactionMessageType
     */
    hTTPTransactionDuration?: DateRangeType;
    /**
     * 
     * @type {ErrorType}
     * @memberof VaultHTTPTransactionMessageType
     */
    hTTPError?: ErrorType;
    /**
     * 
     * @type {VaultHTTPTransactionMessageTypeAuthorizationApproval}
     * @memberof VaultHTTPTransactionMessageType
     */
    authorizationApproval?: VaultHTTPTransactionMessageTypeAuthorizationApproval;
    /**
     * 
     * @type {VaultHTTPTransactionType}
     * @memberof VaultHTTPTransactionMessageType
     */
    type?: VaultHTTPTransactionType;
}

/**
 * Check if a given object implements the VaultHTTPTransactionMessageType interface.
 */
export function instanceOfVaultHTTPTransactionMessageType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function VaultHTTPTransactionMessageTypeFromJSON(json: any): VaultHTTPTransactionMessageType {
    return VaultHTTPTransactionMessageTypeFromJSONTyped(json, false);
}

export function VaultHTTPTransactionMessageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): VaultHTTPTransactionMessageType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'escapedRequestContent': !exists(json, 'escapedRequestContent') ? undefined : json['escapedRequestContent'],
        'escapedResponseContent': !exists(json, 'escapedResponseContent') ? undefined : json['escapedResponseContent'],
        'hTTPTransactionDuration': !exists(json, 'hTTPTransactionDuration') ? undefined : DateRangeTypeFromJSON(json['hTTPTransactionDuration']),
        'hTTPError': !exists(json, 'hTTPError') ? undefined : ErrorTypeFromJSON(json['hTTPError']),
        'authorizationApproval': !exists(json, 'authorizationApproval') ? undefined : VaultHTTPTransactionMessageTypeAuthorizationApprovalFromJSON(json['authorizationApproval']),
        'type': !exists(json, 'type') ? undefined : VaultHTTPTransactionTypeFromJSON(json['type']),
    };
}

export function VaultHTTPTransactionMessageTypeToJSON(value?: VaultHTTPTransactionMessageType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'escapedRequestContent': value.escapedRequestContent,
        'escapedResponseContent': value.escapedResponseContent,
        'hTTPTransactionDuration': DateRangeTypeToJSON(value.hTTPTransactionDuration),
        'hTTPError': ErrorTypeToJSON(value.hTTPError),
        'authorizationApproval': VaultHTTPTransactionMessageTypeAuthorizationApprovalToJSON(value.authorizationApproval),
        'type': VaultHTTPTransactionTypeToJSON(value.type),
    };
}

