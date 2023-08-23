/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AuthorizerTrxLimitType } from './AuthorizerTrxLimitType';
import {
    AuthorizerTrxLimitTypeFromJSON,
    AuthorizerTrxLimitTypeFromJSONTyped,
    AuthorizerTrxLimitTypeToJSON,
} from './AuthorizerTrxLimitType';

/**
 * A representation of the information contained by a Authorizer Group.
 * @export
 * @interface AuthorizerGroupType
 */
export interface AuthorizerGroupType {
    /**
     * Description of the Authorizer Group.
     * @type {string}
     * @memberof AuthorizerGroupType
     */
    description?: string;
    /**
     * Rate Code of the Authorizer Group.
     * @type {string}
     * @memberof AuthorizerGroupType
     */
    rateCode?: string;
    /**
     * Set of Transaction Limit configured for the Authorizer Group.
     * @type {Array<AuthorizerTrxLimitType>}
     * @memberof AuthorizerGroupType
     */
    transactionLimits?: Array<AuthorizerTrxLimitType>;
    /**
     * Hotel Code to which the Authorizer Group belongs to.
     * @type {string}
     * @memberof AuthorizerGroupType
     */
    hotelId?: string;
    /**
     * Authorizer Group code.
     * @type {string}
     * @memberof AuthorizerGroupType
     */
    code?: string;
}

/**
 * Check if a given object implements the AuthorizerGroupType interface.
 */
export function instanceOfAuthorizerGroupType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AuthorizerGroupTypeFromJSON(json: any): AuthorizerGroupType {
    return AuthorizerGroupTypeFromJSONTyped(json, false);
}

export function AuthorizerGroupTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthorizerGroupType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': !exists(json, 'description') ? undefined : json['description'],
        'rateCode': !exists(json, 'rateCode') ? undefined : json['rateCode'],
        'transactionLimits': !exists(json, 'transactionLimits') ? undefined : ((json['transactionLimits'] as Array<any>).map(AuthorizerTrxLimitTypeFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'code': !exists(json, 'code') ? undefined : json['code'],
    };
}

export function AuthorizerGroupTypeToJSON(value?: AuthorizerGroupType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
        'rateCode': value.rateCode,
        'transactionLimits': value.transactionLimits === undefined ? undefined : ((value.transactionLimits as Array<any>).map(AuthorizerTrxLimitTypeToJSON)),
        'hotelId': value.hotelId,
        'code': value.code,
    };
}

