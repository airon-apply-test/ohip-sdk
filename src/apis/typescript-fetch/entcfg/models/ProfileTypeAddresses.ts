/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AddressInfoType } from './AddressInfoType';
import {
    AddressInfoTypeFromJSON,
    AddressInfoTypeFromJSONTyped,
    AddressInfoTypeToJSON,
} from './AddressInfoType';

/**
 * List of customer addresses.
 * @export
 * @interface ProfileTypeAddresses
 */
export interface ProfileTypeAddresses {
    /**
     * Collection of Detailed information on an address for the customer.
     * @type {Array<AddressInfoType>}
     * @memberof ProfileTypeAddresses
     */
    addressInfo?: Array<AddressInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeAddresses
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeAddresses
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeAddresses
     */
    count?: number;
}

/**
 * Check if a given object implements the ProfileTypeAddresses interface.
 */
export function instanceOfProfileTypeAddresses(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileTypeAddressesFromJSON(json: any): ProfileTypeAddresses {
    return ProfileTypeAddressesFromJSONTyped(json, false);
}

export function ProfileTypeAddressesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeAddresses {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addressInfo': !exists(json, 'addressInfo') ? undefined : ((json['addressInfo'] as Array<any>).map(AddressInfoTypeFromJSON)),
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function ProfileTypeAddressesToJSON(value?: ProfileTypeAddresses | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addressInfo': value.addressInfo === undefined ? undefined : ((value.addressInfo as Array<any>).map(AddressInfoTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

