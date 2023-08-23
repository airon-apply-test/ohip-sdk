/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { NameValueModuleType } from './NameValueModuleType';
import {
    NameValueModuleTypeFromJSON,
    NameValueModuleTypeFromJSONTyped,
    NameValueModuleTypeToJSON,
} from './NameValueModuleType';

/**
 * Contains origin details.
 * @export
 * @interface NameValueOriginType
 */
export interface NameValueOriginType {
    /**
     * 
     * @type {NameValueModuleType}
     * @memberof NameValueOriginType
     */
    originName?: NameValueModuleType;
    /**
     * Contains destination column for Origin.
     * @type {string}
     * @memberof NameValueOriginType
     */
    destination?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof NameValueOriginType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof NameValueOriginType
     */
    type?: string;
}

/**
 * Check if a given object implements the NameValueOriginType interface.
 */
export function instanceOfNameValueOriginType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function NameValueOriginTypeFromJSON(json: any): NameValueOriginType {
    return NameValueOriginTypeFromJSONTyped(json, false);
}

export function NameValueOriginTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NameValueOriginType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'originName': !exists(json, 'originName') ? undefined : NameValueModuleTypeFromJSON(json['originName']),
        'destination': !exists(json, 'destination') ? undefined : json['destination'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function NameValueOriginTypeToJSON(value?: NameValueOriginType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'originName': NameValueModuleTypeToJSON(value.originName),
        'destination': value.destination,
        'id': value.id,
        'type': value.type,
    };
}

