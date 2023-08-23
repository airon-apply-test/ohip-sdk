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
import type { ResCommunicationTypeEmails } from './ResCommunicationTypeEmails';
import {
    ResCommunicationTypeEmailsFromJSON,
    ResCommunicationTypeEmailsFromJSONTyped,
    ResCommunicationTypeEmailsToJSON,
} from './ResCommunicationTypeEmails';
import type { ResCommunicationTypeTelephones } from './ResCommunicationTypeTelephones';
import {
    ResCommunicationTypeTelephonesFromJSON,
    ResCommunicationTypeTelephonesFromJSONTyped,
    ResCommunicationTypeTelephonesToJSON,
} from './ResCommunicationTypeTelephones';

/**
 * Communication details for a reservation.
 * @export
 * @interface ResCommunicationType
 */
export interface ResCommunicationType {
    /**
     * 
     * @type {ResCommunicationTypeTelephones}
     * @memberof ResCommunicationType
     */
    telephones?: ResCommunicationTypeTelephones;
    /**
     * 
     * @type {ResCommunicationTypeEmails}
     * @memberof ResCommunicationType
     */
    emails?: ResCommunicationTypeEmails;
}

/**
 * Check if a given object implements the ResCommunicationType interface.
 */
export function instanceOfResCommunicationType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResCommunicationTypeFromJSON(json: any): ResCommunicationType {
    return ResCommunicationTypeFromJSONTyped(json, false);
}

export function ResCommunicationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResCommunicationType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'telephones': !exists(json, 'telephones') ? undefined : ResCommunicationTypeTelephonesFromJSON(json['telephones']),
        'emails': !exists(json, 'emails') ? undefined : ResCommunicationTypeEmailsFromJSON(json['emails']),
    };
}

export function ResCommunicationTypeToJSON(value?: ResCommunicationType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'telephones': ResCommunicationTypeTelephonesToJSON(value.telephones),
        'emails': ResCommunicationTypeEmailsToJSON(value.emails),
    };
}

