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
import type { ARApplyPaymentCriteriaType } from './ARApplyPaymentCriteriaType';
import {
    ARApplyPaymentCriteriaTypeFromJSON,
    ARApplyPaymentCriteriaTypeFromJSONTyped,
    ARApplyPaymentCriteriaTypeToJSON,
} from './ARApplyPaymentCriteriaType';
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
 * @interface CreditToInvoices
 */
export interface CreditToInvoices {
    /**
     * 
     * @type {ARApplyPaymentCriteriaType}
     * @memberof CreditToInvoices
     */
    criteria?: ARApplyPaymentCriteriaType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CreditToInvoices
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CreditToInvoices
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CreditToInvoices interface.
 */
export function instanceOfCreditToInvoices(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreditToInvoicesFromJSON(json: any): CreditToInvoices {
    return CreditToInvoicesFromJSONTyped(json, false);
}

export function CreditToInvoicesFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreditToInvoices {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : ARApplyPaymentCriteriaTypeFromJSON(json['criteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CreditToInvoicesToJSON(value?: CreditToInvoices | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': ARApplyPaymentCriteriaTypeToJSON(value.criteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

