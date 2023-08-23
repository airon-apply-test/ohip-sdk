/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ECertificateScopeType } from './ECertificateScopeType';
import {
    ECertificateScopeTypeFromJSON,
    ECertificateScopeTypeFromJSONTyped,
    ECertificateScopeTypeToJSON,
} from './ECertificateScopeType';

/**
 * Hotel to which certificate is attached to.
 * @export
 * @interface ECertificateInfoTypeHotels
 */
export interface ECertificateInfoTypeHotels {
    /**
     * 
     * @type {Array<string>}
     * @memberof ECertificateInfoTypeHotels
     */
    code?: Array<string>;
    /**
     * 
     * @type {ECertificateScopeType}
     * @memberof ECertificateInfoTypeHotels
     */
    scope?: ECertificateScopeType;
}

/**
 * Check if a given object implements the ECertificateInfoTypeHotels interface.
 */
export function instanceOfECertificateInfoTypeHotels(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ECertificateInfoTypeHotelsFromJSON(json: any): ECertificateInfoTypeHotels {
    return ECertificateInfoTypeHotelsFromJSONTyped(json, false);
}

export function ECertificateInfoTypeHotelsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateInfoTypeHotels {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'scope': !exists(json, 'scope') ? undefined : ECertificateScopeTypeFromJSON(json['scope']),
    };
}

export function ECertificateInfoTypeHotelsToJSON(value?: ECertificateInfoTypeHotels | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'scope': ECertificateScopeTypeToJSON(value.scope),
    };
}

