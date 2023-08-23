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
import type { FinancialPostingsType } from './FinancialPostingsType';
import {
    FinancialPostingsTypeFromJSON,
    FinancialPostingsTypeFromJSONTyped,
    FinancialPostingsTypeToJSON,
} from './FinancialPostingsType';
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
 * Response for fetching the Posting Journal.
 * @export
 * @interface FinancialPostings
 */
export interface FinancialPostings {
    /**
     * 
     * @type {FinancialPostingsType}
     * @memberof FinancialPostings
     */
    journalPostings?: FinancialPostingsType;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof FinancialPostings
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof FinancialPostings
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof FinancialPostings
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof FinancialPostings
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof FinancialPostings
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof FinancialPostings
     */
    count?: number;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof FinancialPostings
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FinancialPostings
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the FinancialPostings interface.
 */
export function instanceOfFinancialPostings(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FinancialPostingsFromJSON(json: any): FinancialPostings {
    return FinancialPostingsFromJSONTyped(json, false);
}

export function FinancialPostingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): FinancialPostings {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'journalPostings': !exists(json, 'journalPostings') ? undefined : FinancialPostingsTypeFromJSON(json['journalPostings']),
        'totalPages': !exists(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !exists(json, 'offset') ? undefined : json['offset'],
        'limit': !exists(json, 'limit') ? undefined : json['limit'],
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function FinancialPostingsToJSON(value?: FinancialPostings | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'journalPostings': FinancialPostingsTypeToJSON(value.journalPostings),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

