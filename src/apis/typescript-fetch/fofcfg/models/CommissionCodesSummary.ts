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
import type { CommissionCodeSummaryInfoType } from './CommissionCodeSummaryInfoType';
import {
    CommissionCodeSummaryInfoTypeFromJSON,
    CommissionCodeSummaryInfoTypeFromJSONTyped,
    CommissionCodeSummaryInfoTypeToJSON,
} from './CommissionCodeSummaryInfoType';
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
 * Response object for fetching commission codes.
 * @export
 * @interface CommissionCodesSummary
 */
export interface CommissionCodesSummary {
    /**
     * Commission code details.
     * @type {Array<CommissionCodeSummaryInfoType>}
     * @memberof CommissionCodesSummary
     */
    commissionCodesSummary?: Array<CommissionCodeSummaryInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CommissionCodesSummary
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CommissionCodesSummary
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CommissionCodesSummary interface.
 */
export function instanceOfCommissionCodesSummary(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CommissionCodesSummaryFromJSON(json: any): CommissionCodesSummary {
    return CommissionCodesSummaryFromJSONTyped(json, false);
}

export function CommissionCodesSummaryFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommissionCodesSummary {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'commissionCodesSummary': !exists(json, 'commissionCodesSummary') ? undefined : ((json['commissionCodesSummary'] as Array<any>).map(CommissionCodeSummaryInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CommissionCodesSummaryToJSON(value?: CommissionCodesSummary | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'commissionCodesSummary': value.commissionCodesSummary === undefined ? undefined : ((value.commissionCodesSummary as Array<any>).map(CommissionCodeSummaryInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

