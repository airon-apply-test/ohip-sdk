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
import type { ARAppliedPaymentType } from './ARAppliedPaymentType';
import {
    ARAppliedPaymentTypeFromJSON,
    ARAppliedPaymentTypeFromJSONTyped,
    ARAppliedPaymentTypeToJSON,
} from './ARAppliedPaymentType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { TrxInfoType } from './TrxInfoType';
import {
    TrxInfoTypeFromJSON,
    TrxInfoTypeFromJSONTyped,
    TrxInfoTypeToJSON,
} from './TrxInfoType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response to the applied payment details,Contains errors or warnings if any.
 * @export
 * @interface InvoiceAppliedPayments
 */
export interface InvoiceAppliedPayments {
    /**
     * Applied Payment record.
     * @type {Array<ARAppliedPaymentType>}
     * @memberof InvoiceAppliedPayments
     */
    details?: Array<ARAppliedPaymentType>;
    /**
     * List of Transaction codes info.
     * @type {Array<TrxInfoType>}
     * @memberof InvoiceAppliedPayments
     */
    trxCodesInfo?: Array<TrxInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof InvoiceAppliedPayments
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof InvoiceAppliedPayments
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the InvoiceAppliedPayments interface.
 */
export function instanceOfInvoiceAppliedPayments(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InvoiceAppliedPaymentsFromJSON(json: any): InvoiceAppliedPayments {
    return InvoiceAppliedPaymentsFromJSONTyped(json, false);
}

export function InvoiceAppliedPaymentsFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvoiceAppliedPayments {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'details': !exists(json, 'details') ? undefined : ((json['details'] as Array<any>).map(ARAppliedPaymentTypeFromJSON)),
        'trxCodesInfo': !exists(json, 'trxCodesInfo') ? undefined : ((json['trxCodesInfo'] as Array<any>).map(TrxInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function InvoiceAppliedPaymentsToJSON(value?: InvoiceAppliedPayments | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'details': value.details === undefined ? undefined : ((value.details as Array<any>).map(ARAppliedPaymentTypeToJSON)),
        'trxCodesInfo': value.trxCodesInfo === undefined ? undefined : ((value.trxCodesInfo as Array<any>).map(TrxInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

