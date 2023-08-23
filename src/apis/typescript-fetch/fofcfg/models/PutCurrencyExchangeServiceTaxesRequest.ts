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
import type { CurrencyExchangeServiceTaxType } from './CurrencyExchangeServiceTaxType';
import {
    CurrencyExchangeServiceTaxTypeFromJSON,
    CurrencyExchangeServiceTaxTypeFromJSONTyped,
    CurrencyExchangeServiceTaxTypeToJSON,
} from './CurrencyExchangeServiceTaxType';
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
 * @interface PutCurrencyExchangeServiceTaxesRequest
 */
export interface PutCurrencyExchangeServiceTaxesRequest {
    /**
     * List of currency exchange service taxes
     * @type {Array<CurrencyExchangeServiceTaxType>}
     * @memberof PutCurrencyExchangeServiceTaxesRequest
     */
    currencyExchangeServiceTaxes?: Array<CurrencyExchangeServiceTaxType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PutCurrencyExchangeServiceTaxesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutCurrencyExchangeServiceTaxesRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PutCurrencyExchangeServiceTaxesRequest interface.
 */
export function instanceOfPutCurrencyExchangeServiceTaxesRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PutCurrencyExchangeServiceTaxesRequestFromJSON(json: any): PutCurrencyExchangeServiceTaxesRequest {
    return PutCurrencyExchangeServiceTaxesRequestFromJSONTyped(json, false);
}

export function PutCurrencyExchangeServiceTaxesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutCurrencyExchangeServiceTaxesRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'currencyExchangeServiceTaxes': !exists(json, 'currencyExchangeServiceTaxes') ? undefined : ((json['currencyExchangeServiceTaxes'] as Array<any>).map(CurrencyExchangeServiceTaxTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PutCurrencyExchangeServiceTaxesRequestToJSON(value?: PutCurrencyExchangeServiceTaxesRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'currencyExchangeServiceTaxes': value.currencyExchangeServiceTaxes === undefined ? undefined : ((value.currencyExchangeServiceTaxes as Array<any>).map(CurrencyExchangeServiceTaxTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

