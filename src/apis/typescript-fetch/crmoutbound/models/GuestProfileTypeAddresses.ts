/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
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
 * @interface GuestProfileTypeAddresses
 */
export interface GuestProfileTypeAddresses {
    /**
     * Collection of Detailed information on an address for the customer.
     * @type {Array<AddressInfoType>}
     * @memberof GuestProfileTypeAddresses
     */
    addressInfo?: Array<AddressInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof GuestProfileTypeAddresses
     */
    allRowsFetched?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof GuestProfileTypeAddresses
     */
    totalRows?: number;
}

/**
 * Check if a given object implements the GuestProfileTypeAddresses interface.
 */
export function instanceOfGuestProfileTypeAddresses(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GuestProfileTypeAddressesFromJSON(json: any): GuestProfileTypeAddresses {
    return GuestProfileTypeAddressesFromJSONTyped(json, false);
}

export function GuestProfileTypeAddressesFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuestProfileTypeAddresses {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addressInfo': !exists(json, 'addressInfo') ? undefined : ((json['addressInfo'] as Array<any>).map(AddressInfoTypeFromJSON)),
        'allRowsFetched': !exists(json, 'allRowsFetched') ? undefined : json['allRowsFetched'],
        'totalRows': !exists(json, 'totalRows') ? undefined : json['totalRows'],
    };
}

export function GuestProfileTypeAddressesToJSON(value?: GuestProfileTypeAddresses | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addressInfo': value.addressInfo === undefined ? undefined : ((value.addressInfo as Array<any>).map(AddressInfoTypeToJSON)),
        'allRowsFetched': value.allRowsFetched,
        'totalRows': value.totalRows,
    };
}

