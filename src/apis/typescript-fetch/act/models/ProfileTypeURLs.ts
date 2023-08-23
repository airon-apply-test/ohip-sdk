/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { URLInfoType } from './URLInfoType';
import {
    URLInfoTypeFromJSON,
    URLInfoTypeFromJSONTyped,
    URLInfoTypeToJSON,
} from './URLInfoType';

/**
 * List of Information on a URL for the customer.
 * @export
 * @interface ProfileTypeURLs
 */
export interface ProfileTypeURLs {
    /**
     * Collection of Detailed information on web url/address for the customer.
     * @type {Array<URLInfoType>}
     * @memberof ProfileTypeURLs
     */
    uRLInfo?: Array<URLInfoType>;
}

/**
 * Check if a given object implements the ProfileTypeURLs interface.
 */
export function instanceOfProfileTypeURLs(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileTypeURLsFromJSON(json: any): ProfileTypeURLs {
    return ProfileTypeURLsFromJSONTyped(json, false);
}

export function ProfileTypeURLsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeURLs {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uRLInfo': !exists(json, 'uRLInfo') ? undefined : ((json['uRLInfo'] as Array<any>).map(URLInfoTypeFromJSON)),
    };
}

export function ProfileTypeURLsToJSON(value?: ProfileTypeURLs | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uRLInfo': value.uRLInfo === undefined ? undefined : ((value.uRLInfo as Array<any>).map(URLInfoTypeToJSON)),
    };
}

