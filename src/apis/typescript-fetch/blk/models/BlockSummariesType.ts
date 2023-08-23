/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BlockSummariesTypeBlockInfoInner } from './BlockSummariesTypeBlockInfoInner';
import {
    BlockSummariesTypeBlockInfoInnerFromJSON,
    BlockSummariesTypeBlockInfoInnerFromJSONTyped,
    BlockSummariesTypeBlockInfoInnerToJSON,
} from './BlockSummariesTypeBlockInfoInner';

/**
 * Summary information returned during block search.
 * @export
 * @interface BlockSummariesType
 */
export interface BlockSummariesType {
    /**
     * A collection of Blocks or Unique IDs of Blocks.
     * @type {Array<BlockSummariesTypeBlockInfoInner>}
     * @memberof BlockSummariesType
     */
    blockInfo?: Array<BlockSummariesTypeBlockInfoInner>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof BlockSummariesType
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof BlockSummariesType
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof BlockSummariesType
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof BlockSummariesType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof BlockSummariesType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof BlockSummariesType
     */
    count?: number;
}

/**
 * Check if a given object implements the BlockSummariesType interface.
 */
export function instanceOfBlockSummariesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockSummariesTypeFromJSON(json: any): BlockSummariesType {
    return BlockSummariesTypeFromJSONTyped(json, false);
}

export function BlockSummariesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockSummariesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockInfo': !exists(json, 'blockInfo') ? undefined : ((json['blockInfo'] as Array<any>).map(BlockSummariesTypeBlockInfoInnerFromJSON)),
        'totalPages': !exists(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !exists(json, 'offset') ? undefined : json['offset'],
        'limit': !exists(json, 'limit') ? undefined : json['limit'],
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function BlockSummariesTypeToJSON(value?: BlockSummariesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockInfo': value.blockInfo === undefined ? undefined : ((value.blockInfo as Array<any>).map(BlockSummariesTypeBlockInfoInnerToJSON)),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

