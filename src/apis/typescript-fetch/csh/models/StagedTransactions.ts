/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response to the request to get staged transaction information for a Reservation
 * @export
 * @interface StagedTransactions
 */
export interface StagedTransactions {
    /**
     * Transaction ID List of staged comp transactions for the reservation.
     * @type {Array<Array<UniqueIDType>>}
     * @memberof StagedTransactions
     */
    stagedCompTransactionsInfo?: Array<Array<UniqueIDType>>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof StagedTransactions
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof StagedTransactions
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the StagedTransactions interface.
 */
export function instanceOfStagedTransactions(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StagedTransactionsFromJSON(json: any): StagedTransactions {
    return StagedTransactionsFromJSONTyped(json, false);
}

export function StagedTransactionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): StagedTransactions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'stagedCompTransactionsInfo': !exists(json, 'stagedCompTransactionsInfo') ? undefined : json['stagedCompTransactionsInfo'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function StagedTransactionsToJSON(value?: StagedTransactions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'stagedCompTransactionsInfo': value.stagedCompTransactionsInfo,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

