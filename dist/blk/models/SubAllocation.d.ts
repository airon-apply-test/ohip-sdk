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
import type { BlockDetailInstructionType } from './BlockDetailInstructionType';
import type { InstanceLink } from './InstanceLink';
import type { SubAllocationsType } from './SubAllocationsType';
import type { WarningType } from './WarningType';
/**
 * Request object for creation of Sub Allocation. This object contains sub allocation details with unique identifiers for each sub allocation along with Master Allocation information. The standard optional Opera Context element is also included.
 * @export
 * @interface SubAllocation
 */
export interface SubAllocation {
    /**
     *
     * @type {SubAllocationsType}
     * @memberof SubAllocation
     */
    subAllocations?: SubAllocationsType;
    /**
     * The instruction to determine the block information to be returned in a successful create operation.
     * @type {Array<BlockDetailInstructionType>}
     * @memberof SubAllocation
     */
    fetchInstructions?: Array<BlockDetailInstructionType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof SubAllocation
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof SubAllocation
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the SubAllocation interface.
 */
export declare function instanceOfSubAllocation(value: object): boolean;
export declare function SubAllocationFromJSON(json: any): SubAllocation;
export declare function SubAllocationFromJSONTyped(json: any, ignoreDiscriminator: boolean): SubAllocation;
export declare function SubAllocationToJSON(value?: SubAllocation | null): any;
