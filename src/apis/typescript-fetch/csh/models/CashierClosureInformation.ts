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
import type { CashierClosureInfoType } from './CashierClosureInfoType';
import {
    CashierClosureInfoTypeFromJSON,
    CashierClosureInfoTypeFromJSONTyped,
    CashierClosureInfoTypeToJSON,
} from './CashierClosureInfoType';
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
 * Response to the request to get financial details for a Posting Activity. This will return the before and after values for the transactions which were changed.
 * @export
 * @interface CashierClosureInformation
 */
export interface CashierClosureInformation {
    /**
     * Information of the Cashier.
     * @type {Array<CashierClosureInfoType>}
     * @memberof CashierClosureInformation
     */
    cashierClosureInfo?: Array<CashierClosureInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CashierClosureInformation
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CashierClosureInformation
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CashierClosureInformation interface.
 */
export function instanceOfCashierClosureInformation(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CashierClosureInformationFromJSON(json: any): CashierClosureInformation {
    return CashierClosureInformationFromJSONTyped(json, false);
}

export function CashierClosureInformationFromJSONTyped(json: any, ignoreDiscriminator: boolean): CashierClosureInformation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cashierClosureInfo': !exists(json, 'cashierClosureInfo') ? undefined : ((json['cashierClosureInfo'] as Array<any>).map(CashierClosureInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CashierClosureInformationToJSON(value?: CashierClosureInformation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cashierClosureInfo': value.cashierClosureInfo === undefined ? undefined : ((value.cashierClosureInfo as Array<any>).map(CashierClosureInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

